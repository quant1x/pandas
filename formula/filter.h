#pragma once
#ifndef PANDAS_FORMULA_FILTER_H
#define PANDAS_FORMULA_FILTER_H 1

#include "support.h"

namespace quant1x::pandas {
    // FILTER 条件为真后把其后 N 个位置置零，按处理后序列继续扫描。
    inline Series filter(const Series &x, std::size_t n) {
        Series out = x;
        for (std::size_t i = 0; i < out.size(); ++i) {
            if (out[i] == 0.0) {
                continue;
            }
            const std::size_t end = std::min(out.size(), i + 1 + n);
            for (std::size_t j = i + 1; j < end; ++j) {
                out[j] = 0.0;
            }
        }
        return out;
    }
}  // namespace quant1x::pandas

#endif  // PANDAS_FORMULA_FILTER_H
