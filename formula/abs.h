#pragma once
#ifndef PANDAS_FORMULA_ABS_H
#define PANDAS_FORMULA_ABS_H 1

#include "support.h"

namespace quant1x::pandas {

    // ABS 逐元素绝对值。
    inline Series abs(const Series &x) {
        Series out(x.size());
        for (std::size_t i = 0; i < x.size(); ++i) {
            out[i] = std::abs(x[i]);
        }
        return out;
    }

}  // namespace quant1x::pandas

#endif  // PANDAS_FORMULA_ABS_H
