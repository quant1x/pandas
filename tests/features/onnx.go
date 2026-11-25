//go:build onnx && cgo

// ONNX 推理（M8，Go 侧）：cgo 绑定 onnxruntime C API。
//
// 需要 onnxruntime SDK（NuGet 的 `Microsoft.ML.OnnxRuntime` 即可），通过环境变量给到：
//
//	CGO_ENABLED=1 \
//	CGO_CFLAGS="-I<sdk>/build/native/include" \
//	CGO_LDFLAGS="-L<sdk>/runtimes/win-x64/native -lonnxruntime" \
//	PATH="<sdk>/runtimes/win-x64/native:$PATH" \
//	go test -tags onnx ./tests/features/ -run TestONNX
//
// 运行时动态库从 `PATH` 找（`go test` 的二进制在临时目录，exe 旁放库不生效）。
package features

/*
#cgo CFLAGS: -O2
#include <onnxruntime_c_api.h>
#include <stdlib.h>
#include <string.h>

#ifdef _WIN32
#include <windows.h>
#else
#include <dlfcn.h>
#endif

static const char* g_error = NULL;
static const OrtApiBase* g_base = NULL;

// or_load 按绝对路径加载 onnxruntime 动态库并取 API 基址；成功返回 0。
static int or_load(const char* dll_path) {
	g_error = NULL;
	if (g_base != NULL) {
		return 0;
	}
	void* handle = NULL;
#ifdef _WIN32
	handle = (void*)LoadLibraryA(dll_path);
#else
	handle = dlopen(dll_path, RTLD_NOW | RTLD_LOCAL);
#endif
	if (handle == NULL) {
		g_error = "加载 onnxruntime 动态库失败";
		return 1;
	}
	typedef const OrtApiBase* (*ort_get_api_base_fn)(void);
#ifdef _WIN32
	ort_get_api_base_fn fn = (ort_get_api_base_fn)GetProcAddress((HMODULE)handle, "OrtGetApiBase");
#else
	ort_get_api_base_fn fn = (ort_get_api_base_fn)dlsym(handle, "OrtGetApiBase");
#endif
	if (fn == NULL) {
		g_error = "找不到 OrtGetApiBase";
		return 2;
	}
	g_base = fn();
	if (g_base == NULL) {
		g_error = "OrtGetApiBase 返回空";
		return 3;
	}
	return 0;
}

static const OrtApi* or_api(void) {
	if (g_base == NULL) {
		return NULL;
	}
	return g_base->GetApi(ORT_API_VERSION);
}

// or_infer 跑一次模型：输入 [rows, 11] 的 double，输出 rows 个 double。
// 成功返回 0；失败返回非 0，原因可用 or_last_error() 取。
static int or_infer(const void* model_data, size_t model_len, double* data, long long rows,
                    double* out) {
	g_error = NULL;
	const OrtApi* api = or_api();
	if (api == NULL) {
		g_error = "GetApi returned NULL (onnxruntime 与头文件版本不匹配?)";
		return 1;
	}

	OrtEnv* env = NULL;
	OrtSessionOptions* options = NULL;
	OrtSession* session = NULL;
	OrtMemoryInfo* memory = NULL;
	OrtValue* input = NULL;
	OrtValue* output = NULL;
	OrtStatus* status = NULL;
	int result = 0;

	if ((status = api->CreateEnv(ORT_LOGGING_LEVEL_WARNING, "go-onnx", &env)) != NULL) {
		g_error = api->GetErrorMessage(status);
		return 2;
	}
	if ((status = api->CreateSessionOptions(&options)) != NULL) {
		g_error = api->GetErrorMessage(status);
		result = 3;
		goto done;
	}
	if ((status = api->CreateSessionFromArray(env, model_data, model_len, options, &session)) != NULL) {
		g_error = api->GetErrorMessage(status);
		result = 4;
		goto done;
	}
	if ((status = api->CreateCpuMemoryInfo(OrtArenaAllocator, OrtMemTypeDefault, &memory)) != NULL) {
		g_error = api->GetErrorMessage(status);
		result = 5;
		goto done;
	}
	{
		const int64_t dims[2] = {rows, 11};
		if ((status = api->CreateTensorWithDataAsOrtValue(memory, data,
		                                                (size_t)rows * 11 * sizeof(double), dims, 2,
		                                                ONNX_TENSOR_ELEMENT_DATA_TYPE_DOUBLE,
		                                                &input)) != NULL) {
			g_error = api->GetErrorMessage(status);
			result = 6;
			goto done;
		}
	}
	{
		const char* input_names[1] = {"features"};
		const char* output_names[1] = {"output"};
		const OrtValue* const inputs[1] = {input};
		if ((status = api->Run(session, NULL, input_names, inputs, 1, output_names, 1,
		                       &output)) != NULL) {
			g_error = api->GetErrorMessage(status);
			result = 7;
			goto done;
		}
	}
	{
		double* values = NULL;
		if ((status = api->GetTensorMutableData(output, (void**)&values)) != NULL) {
			g_error = api->GetErrorMessage(status);
			result = 8;
			goto done;
		}
		memcpy(out, values, (size_t)rows * sizeof(double));
	}

done:
	if (output != NULL) api->ReleaseValue(output);
	if (input != NULL) api->ReleaseValue(input);
	if (memory != NULL) api->ReleaseMemoryInfo(memory);
	if (session != NULL) api->ReleaseSession(session);
	if (options != NULL) api->ReleaseSessionOptions(options);
	api->ReleaseEnv(env);
	return result;
}

static const char* or_last_error(void) {
	return g_error == NULL ? "" : g_error;
}
*/
import "C"

import (
	"fmt"
	"os"
	"unsafe"
)

// LoadRuntime 按绝对路径加载 onnxruntime 动态库（只需加载一次）。
func LoadRuntime(dllPath string) error {
	path := C.CString(dllPath)
	defer C.free(unsafe.Pointer(path))
	if code := C.or_load(path); code != 0 {
		return fmt.Errorf("加载 %s 失败(code=%d): %s", dllPath, code, C.GoString(C.or_last_error()))
	}
	return nil
}

// InferONNX 用 onnxruntime 跑一次 `mlp.onnx`：data 是 [rows, 11] 的行优先数组。
//
// 模型从内存加载（`CreateSessionFromArray`），避免 Windows 下 `ORTCHAR_T` 的宽字符路径转换。
func InferONNX(modelPath string, data []float64, rows int) ([]float64, error) {
	if rows <= 0 || len(data) != rows*11 {
		return nil, fmt.Errorf("输入形状不对：rows=%d, len=%d", rows, len(data))
	}
	model, err := os.ReadFile(modelPath)
	if err != nil {
		return nil, err
	}

	output := make([]float64, rows)
	code := C.or_infer(unsafe.Pointer(&model[0]), C.size_t(len(model)),
		(*C.double)(unsafe.Pointer(&data[0])), C.longlong(rows),
		(*C.double)(unsafe.Pointer(&output[0])))
	if code != 0 {
		return nil, fmt.Errorf("推理失败(code=%d): %s", code, C.GoString(C.or_last_error()))
	}
	return output, nil
}
