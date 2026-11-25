#pragma once
#ifndef PANDAS_FORMULA_MA_H
#define PANDAS_FORMULA_MA_H 1

#include "support.h"

namespace quant1x::pandas {
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
}  // namespace quant1x::pandas

#endif  // PANDAS_FORMULA_MA_H
