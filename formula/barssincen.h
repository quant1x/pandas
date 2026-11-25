#pragma once
#ifndef PANDAS_FORMULA_BARSSINCEN_H
#define PANDAS_FORMULA_BARSSINCEN_H 1

#include "support.h"

namespace quant1x::pandas {
    // BARSSINCEN 窗口内第一次条件成立到当前的周期数；没有成立为 0。
    inline Series barssincen(const Series &condition, std::size_t window) {
        Series out(condition.size(), 0.0);
        if (window == 0 || window > condition.size()) {
            return out;
        }
        for (std::size_t i = window - 1; i < condition.size(); ++i) {
            const std::size_t begin = i + 1 - window;
            for (std::size_t j = begin; j <= i; ++j) {
                if (condition[j] != 0.0) {
                    out[i] = static_cast<double>(i - j);
                    break;
                }
            }
        }
        return out;
    }
}  // namespace quant1x::pandas

#endif  // PANDAS_FORMULA_BARSSINCEN_H
