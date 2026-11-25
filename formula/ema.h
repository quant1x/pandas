#pragma once
#ifndef PANDAS_FORMULA_EMA_H
#define PANDAS_FORMULA_EMA_H 1

#include "support.h"

namespace quant1x::pandas {
    // EMA 指数移动平均：alpha = 2/(N+1)，首值等于源序列首值。
    inline Series ema(const Series &x, std::size_t span) {
        if (x.empty()) {
            return Series{};
        }
        const std::size_t period = span == 0 ? 1 : span;
        const double alpha = 2.0 / (static_cast<double>(period) + 1.0);
        Series out(x.size());
        out[0] = x[0];
        for (std::size_t i = 1; i < x.size(); ++i) {
            out[i] = (1.0 - alpha) * out[i - 1] + alpha * x[i];
        }
        return out;
    }
}  // namespace quant1x::pandas

#endif  // PANDAS_FORMULA_EMA_H
