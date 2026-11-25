#pragma once
#ifndef PANDAS_FORMULA_SUPPORT_H
#define PANDAS_FORMULA_SUPPORT_H 1

// 公式族的公共前置：序列类型、NaN 构造与滚动框架。
// 对应 Go 侧 pandas 的 `Series.Rolling`（不在 formula/ 目录）。

#include <algorithm>
#include <cmath>
#include <cstddef>
#include <cstdint>
#include <limits>
#include <vector>

namespace quant1x::pandas {

    using Series = std::vector<double>;

    inline double nan_value() { return std::numeric_limits<double>::quiet_NaN(); }

    // 滚动框架：完整窗口才计算，窗口不足为 NaN。
    template<typename F>
    inline Series rolling(const Series &x, std::size_t window, F apply) {
        Series out(x.size(), nan_value());
        if (window == 0 || window > x.size()) {
            return out;
        }
        for (std::size_t i = window - 1; i < x.size(); ++i) {
            out[i] = apply(x, i + 1 - window, i + 1);
        }
        return out;
    }

}  // namespace quant1x::pandas

#endif  // PANDAS_FORMULA_SUPPORT_H
