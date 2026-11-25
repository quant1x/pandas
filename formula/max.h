#pragma once
#ifndef PANDAS_FORMULA_MAX_H
#define PANDAS_FORMULA_MAX_H 1

#include "support.h"

namespace quant1x::pandas {
    // MAX 两条序列逐元素取大，按最长对齐，缺位为 NaN。
    inline Series max_series(const Series &x, const Series &y) {
        const std::size_t length = x.size() > y.size() ? x.size() : y.size();
        Series out(length, nan_value());
        for (std::size_t i = 0; i < length; ++i) {
            const double a = i < x.size() ? x[i] : nan_value();
            const double b = i < y.size() ? y[i] : nan_value();
            out[i] = b > a ? b : a;
        }
        return out;
    }
}  // namespace quant1x::pandas

#endif  // PANDAS_FORMULA_MAX_H
