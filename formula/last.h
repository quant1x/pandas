#pragma once
#ifndef PANDAS_FORMULA_LAST_H
#define PANDAS_FORMULA_LAST_H 1

#include "support.h"

namespace quant1x::pandas {
    // LAST 持续存在：A+1 窗口内排除最近 B 根后全部非零；窗口不足 NaN。
    inline Series last(const Series &condition, std::size_t a, std::size_t b) {
        return rolling(condition, a + 1, [b](const Series &v, std::size_t begin, std::size_t end) {
            if (b >= end - begin) {
                return 1.0;
            }
            const std::size_t stop = end - b;
            for (std::size_t i = begin; i < stop; ++i) {
                if (v[i] == 0.0) {
                    return 0.0;
                }
            }
            return 1.0;
        });
    }
}  // namespace quant1x::pandas

#endif  // PANDAS_FORMULA_LAST_H
