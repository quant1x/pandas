#pragma once
#ifndef PANDAS_FORMULA_LLVBARS_H
#define PANDAS_FORMULA_LLVBARS_H 1

#include "support.h"

namespace quant1x::pandas {
    // LLVBARS 最近 N 周期内最低点到当前的周期数；并列取最近的一个。
    inline Series llvbars(const Series &x, std::size_t window) {
        return rolling(x, window, [](const Series &v, std::size_t begin, std::size_t end) {
            std::size_t position = begin;
            for (std::size_t i = begin + 1; i < end; ++i) {
                if (v[i] <= v[position]) {
                    position = i;
                }
            }
            return static_cast<double>(end - 1 - position);
        });
    }
}  // namespace quant1x::pandas

#endif  // PANDAS_FORMULA_LLVBARS_H
