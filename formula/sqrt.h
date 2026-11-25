#pragma once
#ifndef PANDAS_FORMULA_SQRT_H
#define PANDAS_FORMULA_SQRT_H 1

#include "support.h"

namespace quant1x::pandas {
    // SQRT 逐元素平方根。
    inline Series sqrt(const Series &x) {
        Series out(x.size());
        for (std::size_t i = 0; i < x.size(); ++i) {
            out[i] = std::sqrt(x[i]);
        }
        return out;
    }
}  // namespace quant1x::pandas

#endif  // PANDAS_FORMULA_SQRT_H
