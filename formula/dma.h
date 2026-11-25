#pragma once
#ifndef PANDAS_FORMULA_DMA_H
#define PANDAS_FORMULA_DMA_H 1

#include "support.h"

namespace quant1x::pandas {
    // DMA 动态移动平均：alpha = 1/A，A 可常量或逐位置；非有限时按 1。
    inline Series dma(const Series &x, const Series &a) {
        if (x.empty()) {
            return Series{};
        }
        Series out(x.size());
        out[0] = x[0];
        for (std::size_t i = 1; i < x.size(); ++i) {
            const double period = a.size() == 1 ? a[0] : a[i];
            double alpha = 1.0 / period;
            if (!std::isfinite(alpha)) {
                alpha = 1.0;
            }
            out[i] = alpha * x[i] + (1.0 - alpha) * out[i - 1];
        }
        return out;
    }
}  // namespace quant1x::pandas

#endif  // PANDAS_FORMULA_DMA_H
