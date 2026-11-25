#pragma once
#ifndef PANDAS_FORMULA_BARSLASTCOUNT_H
#define PANDAS_FORMULA_BARSLASTCOUNT_H 1

#include "support.h"

namespace quant1x::pandas {
    // BARSLASTCOUNT 连续满足条件的周期数。
    inline Series barslastcount(const Series &condition) {
        Series out(condition.size(), 0.0);
        double streak = 0.0;
        for (std::size_t i = 0; i < condition.size(); ++i) {
            if (condition[i] != 0.0) {
                streak += 1.0;
                out[i] = streak;
            } else {
                streak = 0.0;
            }
        }
        return out;
    }
}  // namespace quant1x::pandas

#endif  // PANDAS_FORMULA_BARSLASTCOUNT_H
