// 读取 tests/formula/cases.csv，与 B2 实现逐值对拍。
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

std::vector<std::string> split(const std::string &text, char delimiter) {
    std::vector<std::string> out;
    std::string item;
    std::istringstream stream(text);
    while (std::getline(stream, item, delimiter)) {
        out.push_back(item);
    }
    return out;
}

Series parse_list(const std::string &text) {
    Series out;
    if (text.empty()) {
        return out;
    }
    for (const auto &item : split(text, ';')) {
        if (item == "nan") {
            out.push_back(std::numeric_limits<double>::quiet_NaN());
        } else {
            out.push_back(std::stod(item));
        }
    }
    return out;
}

std::int64_t parse_optional(const std::string &text) {
    return text.empty() ? 0 : std::stoll(text);
}

Series bools_to_series(const std::vector<bool> &flags) {
    Series out(flags.size(), 0.0);
    for (std::size_t i = 0; i < flags.size(); ++i) {
        out[i] = flags[i] ? 1.0 : 0.0;
    }
    return out;
}

std::vector<bool> to_bools(const Series &values) {
    std::vector<bool> out(values.size(), false);
    for (std::size_t i = 0; i < values.size(); ++i) {
        out[i] = values[i] != 0.0;
    }
    return out;
}

bool close(double actual, double expected) {
    if (std::isnan(actual) && std::isnan(expected)) {
        return true;
    }
    if (actual == expected) {
        return true;
    }
    return std::abs(actual - expected) <= 1e-12 * std::max(1.0, std::abs(expected));
}

}  // namespace

int main() {
    std::ifstream input(PANDAS_FORMULA_FIXTURES);
    if (!input) {
        std::printf("cannot open fixture: %s\n", PANDAS_FORMULA_FIXTURES);
        return 1;
    }
    std::string line;
    std::getline(input, line);
    int failures = 0;
    int count = 0;
    while (std::getline(input, line)) {
        if (line.empty()) {
            continue;
        }
        const auto fields = split(line, ',');
        if (fields.size() != 6) {
            std::printf("bad row: %s\n", line.c_str());
            return 1;
        }
        const std::string name = fields[0];
        const Series x = parse_list(fields[1]);
        const Series y = parse_list(fields[2]);
        const std::int64_t n = parse_optional(fields[3]);
        const std::int64_t m = parse_optional(fields[4]);
        const Series want = parse_list(fields[5]);
        Series got;
        if (name == "abs") {
            got = quant1x::pandas::abs(x);
        } else if (name == "ref") {
            got = quant1x::pandas::ref(x, n);
        } else if (name == "ma") {
            got = quant1x::pandas::ma(x, static_cast<std::size_t>(n));
        } else if (name == "ema") {
            got = quant1x::pandas::ema(x, static_cast<std::size_t>(n));
        } else if (name == "sma") {
            got = quant1x::pandas::sma(x, static_cast<std::size_t>(n), static_cast<int>(m));
        } else if (name == "hhv") {
            got = quant1x::pandas::hhv(x, static_cast<std::size_t>(n));
        } else if (name == "llv") {
            got = quant1x::pandas::llv(x, static_cast<std::size_t>(n));
        } else if (name == "sum") {
            got = quant1x::pandas::sum(x, static_cast<std::size_t>(n));
        } else if (name == "std") {
            got = quant1x::pandas::stddev(x, static_cast<std::size_t>(n));
        } else if (name == "wma") {
            got = quant1x::pandas::wma(x, static_cast<std::size_t>(n));
        } else if (name == "sqrt") {
            got = quant1x::pandas::sqrt(x);
        } else if (name == "count") {
            got = quant1x::pandas::count(x, static_cast<std::size_t>(n));
        } else if (name == "filter") {
            got = quant1x::pandas::filter(x, static_cast<std::size_t>(n));
        } else if (name == "hhvbars") {
            got = quant1x::pandas::hhvbars(x, static_cast<std::size_t>(n));
        } else if (name == "llvbars") {
            got = quant1x::pandas::llvbars(x, static_cast<std::size_t>(n));
        } else if (name == "cross") {
            const auto flags = quant1x::pandas::cross(x, y);
            got.assign(flags.size(), 0.0);
            for (std::size_t i = 0; i < flags.size(); ++i) {
                got[i] = flags[i] ? 1.0 : 0.0;
            }
        } else if (name == "dma") {
            got = quant1x::pandas::dma(x, y);
        } else if (name == "max") {
            got = quant1x::pandas::max_series(x, y);
        } else if (name == "min") {
            got = quant1x::pandas::min_series(x, y);
        } else if (name == "gt") {
            got = bools_to_series(quant1x::pandas::gt(x, y));
        } else if (name == "gte") {
            got = bools_to_series(quant1x::pandas::gte(x, y));
        } else if (name == "lt") {
            got = bools_to_series(quant1x::pandas::lt(x, y));
        } else if (name == "lte") {
            got = bools_to_series(quant1x::pandas::lte(x, y));
        } else if (name == "eq") {
            got = bools_to_series(quant1x::pandas::eq(x, y));
        } else if (name == "neq") {
            got = bools_to_series(quant1x::pandas::neq(x, y));
        } else if (name == "and") {
            got = bools_to_series(quant1x::pandas::logic_and(to_bools(x), to_bools(y)));
        } else if (name == "or") {
            got = bools_to_series(quant1x::pandas::logic_or(to_bools(x), to_bools(y)));
        } else if (name == "not") {
            got = bools_to_series(quant1x::pandas::logic_not(to_bools(x)));
        } else if (name == "barslast") {
            got = quant1x::pandas::barslast(x);
        } else if (name == "barslastcount") {
            got = quant1x::pandas::barslastcount(x);
        } else if (name == "barslasts") {
            got = quant1x::pandas::barslasts(x, n);
        } else if (name == "barssincen") {
            got = quant1x::pandas::barssincen(x, static_cast<std::size_t>(n));
        } else if (name == "if") {
            const Series fallback(y.size(), static_cast<double>(n));
            got = quant1x::pandas::if_select(x, y, fallback);
        } else if (name == "ifn") {
            const Series fallback(y.size(), static_cast<double>(n));
            got = quant1x::pandas::ifn(x, y, fallback);
        } else if (name == "last") {
            got = quant1x::pandas::last(x, static_cast<std::size_t>(n), static_cast<std::size_t>(m));
        } else {
            std::printf("unknown formula: %s\n", name.c_str());
            return 1;
        }
        if (got.size() != want.size()) {
            std::printf("%s: size mismatch\n", name.c_str());
            ++failures;
            continue;
        }
        for (std::size_t i = 0; i < want.size(); ++i) {
            if (!close(got[i], want[i])) {
                std::printf("%s: index %zu: %g != %g\n", name.c_str(), i, got[i], want[i]);
                ++failures;
                break;
            }
        }
        ++count;
    }
    if (count < 35 || failures != 0) {
        std::printf("fixture failed: count=%d failures=%d\n", count, failures);
        return 1;
    }
    std::printf("formula fixtures passed (%d cases)\n", count);
    return 0;
}
