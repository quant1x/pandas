#pragma once
#ifndef PANDAS_FORMULA_SUM_H
#define PANDAS_FORMULA_SUM_H 1

#include "support.h"

namespace quant1x::pandas {
    // SUM 最近 N 个周期的和。
    inline Series sum(const Series &x, std::size_t window) {
        return rolling(x, window, [](const Series &v, std::size_t begin, std::size_t end) {
            double total = 0.0;
            for (std::size_t i = begin; i < end; ++i) {
                total += v[i];
            }
            return total;
        });
    }
}  // namespace quant1x::pandas

#endif  // PANDAS_FORMULA_SUM_H
