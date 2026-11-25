// ONNX 推理对拍（M8）：用 onnxruntime C API 加载 `mlp.onnx`，跑 `features.csv` 的特征并与 `pred_mlp` 比较。
//
// 需要在配置期给到 SDK（NuGet 的 `Microsoft.ML.OnnxRuntime` 包即可）：
//   -DPANDAS_WITH_ONNX=ON -DPANDAS_ORT_INCLUDE_DIR=... -DPANDAS_ORT_LIB=... -DPANDAS_ORT_DLL=...
#include <onnxruntime_c_api.h>

#include <cmath>
#include <cstdio>
#include <cstdlib>
#include <fstream>
#include <limits>
#include <sstream>
#include <string>
#include <vector>

namespace {

constexpr const char *kFeatures[11] = {"ma5", "ma20", "ema12", "wma10", "hhv20", "llv20",
                                      "std20", "hhvbars20", "ret1", "bias5", "cross_up"};

std::vector<std::string> split(const std::string &text, char delimiter) {
    std::vector<std::string> out;
    std::string item;
    std::istringstream stream(text);
    while (std::getline(stream, item, delimiter)) {
        out.push_back(item);
    }
    return out;
}

double parse_value(const std::string &text) {
    if (text == "nan") {
        return std::numeric_limits<double>::quiet_NaN();
    }
    return std::stod(text);
}

std::wstring widen(const std::string &text) {
    return std::wstring(text.begin(), text.end());
}

bool check(OrtStatus *status, const char *what, const OrtApi *api) {
    if (status == nullptr) {
        return true;
    }
    std::printf("%s: %s\n", what, api->GetErrorMessage(status));
    api->ReleaseStatus(status);
    return false;
}

}  // namespace

int main() {
    std::ifstream input(PANDAS_FEATURE_FIXTURES);
    if (!input) {
        std::printf("cannot open fixture: %s\n", PANDAS_FEATURE_FIXTURES);
        return 1;
    }

    std::string line;
    std::getline(input, line);
    const auto header = split(line, ',');
    auto index_of = [&header](const char *name) -> std::size_t {
        for (std::size_t i = 0; i < header.size(); ++i) {
            if (header[i] == name) {
                return i;
            }
        }
        std::printf("missing column: %s\n", name);
        std::exit(1);
    };

    std::vector<double> inputs;
    std::vector<double> expected;
    std::int64_t rows = 0;
    while (std::getline(input, line)) {
        if (line.empty()) {
            continue;
        }
        const auto fields = split(line, ',');
        for (const char *name : kFeatures) {
            inputs.push_back(parse_value(fields[index_of(name)]));
        }
        expected.push_back(parse_value(fields[index_of("pred_mlp")]));
        ++rows;
    }
    if (rows < 20) {
        std::printf("too few rows: %lld\n", static_cast<long long>(rows));
        return 1;
    }

    const OrtApiBase *base = OrtGetApiBase();
    const OrtApi *api = base->GetApi(ORT_API_VERSION);
    if (api == nullptr) {
        std::printf("OrtGetApiBase()->GetApi 返回空\n");
        return 1;
    }

    OrtEnv *env = nullptr;
    OrtSessionOptions *options = nullptr;
    OrtSession *session = nullptr;
    OrtMemoryInfo *memory = nullptr;
    OrtValue *tensor = nullptr;
    OrtValue *output = nullptr;

    bool ok = check(api->CreateEnv(ORT_LOGGING_LEVEL_WARNING, "pandas-onnx", &env), "CreateEnv", api) &&
              check(api->CreateSessionOptions(&options), "CreateSessionOptions", api) &&
              check(api->CreateSession(env, widen(PANDAS_ONNX_MODEL).c_str(), options, &session),
                    "CreateSession", api) &&
              check(api->CreateCpuMemoryInfo(OrtArenaAllocator, OrtMemTypeDefault, &memory),
                    "CreateCpuMemoryInfo", api);

    const std::int64_t dims[2] = {rows, 11};
    if (ok) {
        ok = check(api->CreateTensorWithDataAsOrtValue(memory, inputs.data(),
                                                      inputs.size() * sizeof(double), dims, 2,
                                                      ONNX_TENSOR_ELEMENT_DATA_TYPE_DOUBLE, &tensor),
                   "CreateTensor", api);
    }

    const char *input_names[1] = {"features"};
    const char *output_names[1] = {"output"};
    if (ok) {
        ok = check(api->Run(session, nullptr, input_names, &tensor, 1, output_names, 1, &output),
                   "Run", api);
    }

    double *data = nullptr;
    if (ok) {
        ok = check(api->GetTensorMutableData(output, reinterpret_cast<void **>(&data)),
                   "GetTensorMutableData", api);
    }

    int failures = 0;
    if (ok) {
        for (std::int64_t i = 0; i < rows; ++i) {
            const double want = expected[static_cast<std::size_t>(i)];
            const double got = data[i];
            if (std::isnan(got) && std::isnan(want)) {
                continue;
            }
            const double tolerance = 1e-12 * std::max(1.0, std::abs(want));
            if (std::abs(got - want) > tolerance) {
                std::printf("index %lld: %g != %g\n", static_cast<long long>(i), got, want);
                ++failures;
                break;
            }
        }
    }

    if (output != nullptr) api->ReleaseValue(output);
    if (tensor != nullptr) api->ReleaseValue(tensor);
    if (memory != nullptr) api->ReleaseMemoryInfo(memory);
    if (session != nullptr) api->ReleaseSession(session);
    if (options != nullptr) api->ReleaseSessionOptions(options);
    if (env != nullptr) api->ReleaseEnv(env);

    if (!ok || failures != 0) {
        return 1;
    }
    std::printf("onnx fixtures passed (%lld rows)\n", static_cast<long long>(rows));
    return 0;
}
