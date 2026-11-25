#pragma once
#ifndef PANDAS_FORMULA_SMA_H
#define PANDAS_FORMULA_SMA_H 1

#include "support.h"

namespace quant1x::pandas {
    // SMA 中国式平滑：alpha = M/N（M=0 视为 1）。
    inline Series sma(const Series &x, std::size_t n, int m) {
        if (x.empty()) {
            return Series{};
        }
        const std::size_t period = n == 0 ? 1 : n;
        const int weight = m == 0 ? 1 : m;
        const double alpha = static_cast<double>(weight) / static_cast<double>(period);
        Series out(x.size());
        out[0] = x[0];
        for (std::size_t i = 1; i < x.size(); ++i) {
            out[i] = (1.0 - alpha) * out[i - 1] + alpha * x[i];
        }
        return out;
    }
}  // namespace quant1x::pandas

#endif  // PANDAS_FORMULA_SMA_H
