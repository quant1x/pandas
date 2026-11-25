// 读取 tests/features/features.csv，用 quant1x_pandas_cpp 重算特征列并逐值对拍。
#include "formula/formula.h"

#include <algorithm>
#include <cmath>
#include <cstdio>
#include <fstream>
#include <sstream>
#include <string>
#include <vector>

namespace {

using quant1x::pandas::Series;

struct Table {
    std::vector<double> close;
    std::vector<double> high;
    std::vector<double> low;
    std::vector<double> ma5;
    std::vector<double> ma20;
    std::vector<double> ema12;
    std::vector<double> wma10;
    std::vector<double> hhv20;
    std::vector<double> llv20;
    std::vector<double> std20;
    std::vector<double> hhvbars20;
    std::vector<double> ret1;
    std::vector<double> bias5;
    std::vector<double> cross_up;
};

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

bool close_enough(double actual, double expected) {
    if (std::isnan(actual) && std::isnan(expected)) {
        return true;
    }
    if (actual == expected) {
        return true;
    }
    return std::abs(actual - expected) <= 1e-12 * std::max(1.0, std::abs(expected));
}

bool expect_column(const char *name, const std::vector<double> &got, const std::vector<double> &want) {
    if (got.size() != want.size()) {
        std::printf("%s: size %zu != %zu\n", name, got.size(), want.size());
        return false;
    }
    for (std::size_t i = 0; i < want.size(); ++i) {
        if (!close_enough(got[i], want[i])) {
            std::printf("%s: index %zu: %g != %g\n", name, i, got[i], want[i]);
            return false;
        }
    }
    return true;
}

}  // namespace

int main() {
    std::ifstream input(PANDAS_FEATURE_FIXTURES);
    if (!input) {
        std::printf("cannot open fixture: %s\n", PANDAS_FEATURE_FIXTURES);
        return 1;
    }

    std::string line;
    if (!std::getline(input, line)) {
        std::printf("empty fixture\n");
        return 1;
    }
    const auto header = split(line, ',');
    std::size_t date_col = 0;
    std::size_t open_col = 0;
    std::size_t high_col = 0;
    std::size_t low_col = 0;
    std::size_t close_col = 0;
    std::size_t ma5_col = 0;
    std::size_t ma20_col = 0;
    std::size_t ema12_col = 0;
    std::size_t wma10_col = 0;
    std::size_t hhv20_col = 0;
    std::size_t llv20_col = 0;
    std::size_t std20_col = 0;
    std::size_t hhvbars_col = 0;
    std::size_t ret1_col = 0;
    std::size_t bias5_col = 0;
    std::size_t cross_col = 0;
    for (std::size_t i = 0; i < header.size(); ++i) {
        const std::string &name = header[i];
        if (name == "date") date_col = i;
        else if (name == "open") open_col = i;
        else if (name == "high") high_col = i;
        else if (name == "low") low_col = i;
        else if (name == "close") close_col = i;
        else if (name == "ma5") ma5_col = i;
        else if (name == "ma20") ma20_col = i;
        else if (name == "ema12") ema12_col = i;
        else if (name == "wma10") wma10_col = i;
        else if (name == "hhv20") hhv20_col = i;
        else if (name == "llv20") llv20_col = i;
        else if (name == "std20") std20_col = i;
        else if (name == "hhvbars20") hhvbars_col = i;
        else if (name == "ret1") ret1_col = i;
        else if (name == "bias5") bias5_col = i;
        else if (name == "cross_up") cross_col = i;
    }
    (void)date_col;
    (void)open_col;

    Table data;
    while (std::getline(input, line)) {
        if (line.empty()) {
            continue;
        }
        const auto fields = split(line, ',');
        data.close.push_back(parse_value(fields[close_col]));
        data.high.push_back(parse_value(fields[high_col]));
        data.low.push_back(parse_value(fields[low_col]));
        data.ma5.push_back(parse_value(fields[ma5_col]));
        data.ma20.push_back(parse_value(fields[ma20_col]));
        data.ema12.push_back(parse_value(fields[ema12_col]));
        data.wma10.push_back(parse_value(fields[wma10_col]));
        data.hhv20.push_back(parse_value(fields[hhv20_col]));
        data.llv20.push_back(parse_value(fields[llv20_col]));
        data.std20.push_back(parse_value(fields[std20_col]));
        data.hhvbars20.push_back(parse_value(fields[hhvbars_col]));
        data.ret1.push_back(parse_value(fields[ret1_col]));
        data.bias5.push_back(parse_value(fields[bias5_col]));
        data.cross_up.push_back(parse_value(fields[cross_col]));
    }

    const Series close = data.close;
    const Series ma5 = quant1x::pandas::ma(close, 5);
    const Series ma20 = quant1x::pandas::ma(close, 20);
    const Series ema12 = quant1x::pandas::ema(close, 12);
    const Series wma10 = quant1x::pandas::wma(close, 10);
    const Series hhv20 = quant1x::pandas::hhv(data.high, 20);
    const Series llv20 = quant1x::pandas::llv(data.low, 20);
    const Series std20 = quant1x::pandas::stddev(close, 20);
    const Series hhvbars20 = quant1x::pandas::hhvbars(data.high, 20);
    const auto cross = quant1x::pandas::cross(ma5, ma20);

    Series ret1(close.size(), std::numeric_limits<double>::quiet_NaN());
    for (std::size_t i = 1; i < close.size(); ++i) {
        ret1[i] = close[i] / close[i - 1] - 1.0;
    }
    Series bias5(close.size(), std::numeric_limits<double>::quiet_NaN());
    for (std::size_t i = 0; i < close.size(); ++i) {
        if (!std::isnan(ma5[i])) {
            bias5[i] = close[i] / ma5[i] - 1.0;
        }
    }
    Series cross_up(cross.size(), 0.0);
    for (std::size_t i = 0; i < cross.size(); ++i) {
        cross_up[i] = cross[i] ? 1.0 : 0.0;
    }

    const bool ok =
        expect_column("ma5", ma5, data.ma5) &&
        expect_column("ma20", ma20, data.ma20) &&
        expect_column("ema12", ema12, data.ema12) &&
        expect_column("wma10", wma10, data.wma10) &&
        expect_column("hhv20", hhv20, data.hhv20) &&
        expect_column("llv20", llv20, data.llv20) &&
        expect_column("std20", std20, data.std20) &&
        expect_column("hhvbars20", hhvbars20, data.hhvbars20) &&
        expect_column("ret1", ret1, data.ret1) &&
        expect_column("bias5", bias5, data.bias5) &&
        expect_column("cross_up", cross_up, data.cross_up);
    if (!ok) {
        return 1;
    }
    std::printf("feature fixtures passed (%zu bars)\n", close.size());
    return 0;
}
