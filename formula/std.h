#pragma once
#ifndef PANDAS_FORMULA_STD_H
#define PANDAS_FORMULA_STD_H 1

#include "support.h"

namespace quant1x::pandas {
    // STD（C++ 名 stddev）最近 N 个周期的总体标准差。
    inline Series stddev(const Series &x, std::size_t window) {
        return rolling(x, window, [](const Series &v, std::size_t begin, std::size_t end) {
            const double n = static_cast<double>(end - begin);
            double mean = 0.0;
            for (std::size_t i = begin; i < end; ++i) {
                mean += v[i];
            }
            mean /= n;
            double total = 0.0;
            for (std::size_t i = begin; i < end; ++i) {
                total += (v[i] - mean) * (v[i] - mean);
            }
            return std::sqrt(total / n);
        });
    }
}  // namespace quant1x::pandas

#endif  // PANDAS_FORMULA_STD_H
