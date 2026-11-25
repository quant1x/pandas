#pragma once
#ifndef PANDAS_FORMULA_WMA_H
#define PANDAS_FORMULA_WMA_H 1

#include "support.h"

namespace quant1x::pandas {
    // WMA 加权移动平均：窗口内按 1..N 递增加权。
    inline Series wma(const Series &x, std::size_t window) {
        return rolling(x, window, [](const Series &v, std::size_t begin, std::size_t end) {
            double total = 0.0;
            double weights = 0.0;
            for (std::size_t i = begin; i < end; ++i) {
                const double weight = static_cast<double>(i - begin + 1);
                total += weight * v[i];
                weights += weight;
            }
            return total / weights;
        });
    }
}  // namespace quant1x::pandas

#endif  // PANDAS_FORMULA_WMA_H
