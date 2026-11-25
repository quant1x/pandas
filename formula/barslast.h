#pragma once
#ifndef PANDAS_FORMULA_BARSLAST_H
#define PANDAS_FORMULA_BARSLAST_H 1

#include "support.h"

namespace quant1x::pandas {
    // BARSLAST 上一次条件成立到当前的周期数；从未成立为 -1。
    inline Series barslast(const Series &condition) {
        Series out(condition.size(), -1.0);
        std::int64_t last_true = -1;
        for (std::size_t i = 0; i < condition.size(); ++i) {
            if (condition[i] != 0.0) {
                last_true = static_cast<std::int64_t>(i);
                out[i] = 0.0;
            } else if (last_true >= 0) {
                out[i] = static_cast<double>(static_cast<std::int64_t>(i) - last_true);
            }
        }
        return out;
    }
}  // namespace quant1x::pandas

#endif  // PANDAS_FORMULA_BARSLAST_H
