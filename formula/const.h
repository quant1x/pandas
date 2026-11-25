#pragma once
#ifndef PANDAS_FORMULA_CONST_H
#define PANDAS_FORMULA_CONST_H 1

#include "support.h"

namespace quant1x::pandas {
    // CONST 取序列最后值为常量序列；Go 先转 float32，故按 f32 精度落值。
    inline Series const_series(const Series &x) {
        Series out(x.size(), nan_value());
        if (x.empty()) {
            return out;
        }
        const double last = static_cast<double>(static_cast<float>(x.back()));
        for (std::size_t i = 0; i < out.size(); ++i) {
            out[i] = last;
        }
        return out;
    }
}  // namespace quant1x::pandas

#endif  // PANDAS_FORMULA_CONST_H
