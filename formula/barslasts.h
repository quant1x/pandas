#pragma once
#ifndef PANDAS_FORMULA_BARSLASTS_H
#define PANDAS_FORMULA_BARSLASTS_H 1

#include "support.h"

namespace quant1x::pandas {
    // BARSLASTS 倒数第 N 次成立到当前的周期数（Go 递推口径）。
    inline Series barslasts(const Series &condition, std::int64_t n) {
        Series current(condition.size(), 0.0);
        double previous = 0.0;
        for (std::size_t i = 0; i < condition.size(); ++i) {
            previous = condition[i] != 0.0 ? 0.0 : previous + 1.0;
            current[i] = previous;
        }
        for (std::int64_t iteration = 1; iteration < n; ++iteration) {
            Series next(current.size(), nan_value());
            for (std::size_t i = 0; i < current.size(); ++i) {
                if (!std::isfinite(current[i])) {
                    continue;
                }
                const double step = current[i] + 1.0;
                const double source = static_cast<double>(i) - step;
                if (source < 0.0) {
                    continue;
                }
                const std::size_t index = static_cast<std::size_t>(source);
                if (index >= current.size() || !std::isfinite(current[index])) {
                    continue;
                }
                next[i] = current[index] + step;
            }
            current = next;
        }
        return current;
    }
}  // namespace quant1x::pandas

#endif  // PANDAS_FORMULA_BARSLASTS_H
