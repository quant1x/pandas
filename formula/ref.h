#pragma once
#ifndef PANDAS_FORMULA_REF_H
#define PANDAS_FORMULA_REF_H 1

#include "support.h"

namespace quant1x::pandas {
    // REF 引用前 N 个周期的值；N 为负时向后引用，越界为 NaN。
    inline Series ref(const Series &x, std::int64_t n) {
        Series out(x.size(), nan_value());
        if (n == 0) {
            return x;
        }
        if (n > 0) {
            for (std::size_t i = static_cast<std::size_t>(n); i < x.size(); ++i) {
                out[i] = x[i - static_cast<std::size_t>(n)];
            }
        } else {
            const std::size_t offset = static_cast<std::size_t>(-n);
            for (std::size_t i = 0; i + offset < x.size(); ++i) {
                out[i] = x[i + offset];
            }
        }
        return out;
    }
}  // namespace quant1x::pandas

#endif  // PANDAS_FORMULA_REF_H
