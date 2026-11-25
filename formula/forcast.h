#pragma once
#ifndef PANDAS_FORMULA_FORCAST_H
#define PANDAS_FORMULA_FORCAST_H 1

#include "support.h"
#include "slope.h"

namespace quant1x::pandas {
    // FORCAST 窗口内线性回归在最后一根上的预测值。
    inline Series forcast(const Series &x, std::size_t window) {
        return rolling(x, window, [](const Series &v, std::size_t begin, std::size_t end) {
            double slope_value = nan_value();
            double intercept = nan_value();
            linear_regression(v, begin, end, slope_value, intercept);
            if (std::isnan(slope_value)) {
                return nan_value();
            }
            const double last = static_cast<double>(end - begin - 1);
            return intercept + slope_value * last;
        });
    }
}  // namespace quant1x::pandas

#endif  // PANDAS_FORMULA_FORCAST_H
