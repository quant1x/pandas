#pragma once
#ifndef PANDAS_FORMULA_SLOPE_H
#define PANDAS_FORMULA_SLOPE_H 1

#include "support.h"

namespace quant1x::pandas {
    // 最小二乘直线：t 取 0..n-1，返回斜率与截距；不足两点为 NaN。
    inline void linear_regression(const Series &v, std::size_t begin, std::size_t end,
                                  double &slope, double &intercept) {
        const std::size_t n = end - begin;
        if (n < 2) {
            slope = nan_value();
            intercept = nan_value();
            return;
        }
        const double count = static_cast<double>(n);
        const double mean_t = static_cast<double>(n - 1) / 2.0;
        double mean_y = 0.0;
        for (std::size_t i = begin; i < end; ++i) {
            mean_y += v[i];
        }
        mean_y /= count;
        double numerator = 0.0;
        double denominator = 0.0;
        for (std::size_t k = 0; k < n; ++k) {
            const double t = static_cast<double>(k) - mean_t;
            numerator += t * (v[begin + k] - mean_y);
            denominator += t * t;
        }
        slope = numerator / denominator;
        intercept = mean_y - slope * mean_t;
    }
    // SLOPE 窗口内线性回归斜率；不足两点为 NaN。
    inline Series slope(const Series &x, std::size_t window) {
        return rolling(x, window, [](const Series &v, std::size_t begin, std::size_t end) {
            double slope_value = nan_value();
            double intercept = nan_value();
            linear_regression(v, begin, end, slope_value, intercept);
            return slope_value;
        });
    }
}  // namespace quant1x::pandas

#endif  // PANDAS_FORMULA_SLOPE_H
