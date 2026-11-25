#pragma once
#ifndef PANDAS_FORMULA_LLV_H
#define PANDAS_FORMULA_LLV_H 1

#include "support.h"

namespace quant1x::pandas {
    // LLV 最近 N 个周期的最小值。
    inline Series llv(const Series &x, std::size_t window) {
        return rolling(x, window, [](const Series &v, std::size_t begin, std::size_t end) {
            double value = v[begin];
            for (std::size_t i = begin + 1; i < end; ++i) {
                if (v[i] < value) {
                    value = v[i];
                }
            }
            return value;
        });
    }
}  // namespace quant1x::pandas

#endif  // PANDAS_FORMULA_LLV_H
