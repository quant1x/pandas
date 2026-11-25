#pragma once
#ifndef PANDAS_FORMULA_H
#define PANDAS_FORMULA_H 1

// pandas 公式族的 C++ 对齐实现（B2）首批：ABS / REF / MA / EMA / SMA。
//
// 序列用 std::vector<double>，与 B1/B3 的最小参数集一致：
//   - REF 越界为 NaN；
//   - MA 只有完整窗口才计算，窗口不足为 NaN；
//   - EMA alpha=2/(N+1)、SMA alpha=M/N（M=0 视为 1），均按 EWM Adjust=false。

#include <cmath>
#include <cstddef>
#include <cstdint>
#include <limits>
#include <vector>

namespace quant1x::pandas {

    using Series = std::vector<double>;

    inline double nan_value() { return std::numeric_limits<double>::quiet_NaN(); }

    // ABS 逐元素绝对值。
    inline Series abs(const Series &x) {
        Series out(x.size());
        for (std::size_t i = 0; i < x.size(); ++i) {
            out[i] = std::abs(x[i]);
        }
        return out;
    }

    // REF 引用前 N 个周期的值；N 为负时向后引用，越界为 NaN。
    inline Series ref(const Series &x, std::int64_t n) {
        Series out(x.size(), nan_value());
        if (n == 0) {
            return x;
        }
        if (n > 0) {
            for (std::size_t i = static_cast<std::size_t>(n); i < x.size(); ++i) {
                out[i] = x[i - static_cast<std::size_t>(n)];
            }
        } else {
            const std::size_t offset = static_cast<std::size_t>(-n);
            for (std::size_t i = 0; i + offset < x.size(); ++i) {
                out[i] = x[i + offset];
            }
        }
        return out;
    }

    // MA 简单移动平均：完整窗口才计算，窗口不足为 NaN。
    inline Series ma(const Series &x, std::size_t window) {
        Series out(x.size(), nan_value());
        if (window == 0 || window > x.size()) {
            return out;
        }
        double total = 0.0;
        for (std::size_t i = 0; i < x.size(); ++i) {
            total += x[i];
            if (i >= window) {
                total -= x[i - window];
            }
            if (i + 1 >= window) {
                out[i] = total / static_cast<double>(window);
            }
        }
        return out;
    }

    // EMA 指数移动平均：alpha = 2/(N+1)，首值等于源序列首值。
    inline Series ema(const Series &x, std::size_t span) {
        if (x.empty()) {
            return Series{};
        }
        const std::size_t period = span == 0 ? 1 : span;
        const double alpha = 2.0 / (static_cast<double>(period) + 1.0);
        Series out(x.size());
        out[0] = x[0];
        for (std::size_t i = 1; i < x.size(); ++i) {
            out[i] = (1.0 - alpha) * out[i - 1] + alpha * x[i];
        }
        return out;
    }

    // SMA 中国式平滑：alpha = M/N（M=0 视为 1）。
    inline Series sma(const Series &x, std::size_t n, int m) {
        if (x.empty()) {
            return Series{};
        }
        const std::size_t period = n == 0 ? 1 : n;
        const int weight = m == 0 ? 1 : m;
        const double alpha = static_cast<double>(weight) / static_cast<double>(period);
        Series out(x.size());
        out[0] = x[0];
        for (std::size_t i = 1; i < x.size(); ++i) {
            out[i] = (1.0 - alpha) * out[i - 1] + alpha * x[i];
        }
        return out;
    }

}  // namespace quant1x::pandas

#endif  // PANDAS_FORMULA_H
