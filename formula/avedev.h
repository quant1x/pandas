#pragma once
#ifndef PANDAS_FORMULA_AVEDEV_H
#define PANDAS_FORMULA_AVEDEV_H 1

#include "support.h"

namespace quant1x::pandas {
    // AVEDEV 平均绝对偏差：窗口内「与均值之差的绝对值」的平均。
    inline Series avedev(const Series &x, std::size_t window) {
        return rolling(x, window, [](const Series &v, std::size_t begin, std::size_t end) {
            const double count = static_cast<double>(end - begin);
            double mean = 0.0;
            for (std::size_t i = begin; i < end; ++i) {
                mean += v[i];
            }
            mean /= count;
            double total = 0.0;
            for (std::size_t i = begin; i < end; ++i) {
                total += std::abs(v[i] - mean);
            }
            return total / count;
        });
    }
}  // namespace quant1x::pandas

#endif  // PANDAS_FORMULA_AVEDEV_H
