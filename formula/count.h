#pragma once
#ifndef PANDAS_FORMULA_COUNT_H
#define PANDAS_FORMULA_COUNT_H 1

#include "support.h"

namespace quant1x::pandas {
    // COUNT 最近 N 个周期内非零条件的个数；窗口不足计 0。
    inline Series count(const Series &condition, std::size_t window) {
        Series out(condition.size(), 0.0);
        if (window == 0 || window > condition.size()) {
            return out;
        }
        for (std::size_t i = window - 1; i < condition.size(); ++i) {
            std::size_t total = 0;
            for (std::size_t j = i + 1 - window; j <= i; ++j) {
                if (condition[j] != 0.0) {
                    ++total;
                }
            }
            out[i] = static_cast<double>(total);
        }
        return out;
    }
}  // namespace quant1x::pandas

#endif  // PANDAS_FORMULA_COUNT_H
