#pragma once
#ifndef PANDAS_FORMULA_CROSS_H
#define PANDAS_FORMULA_CROSS_H 1

#include "support.h"

namespace quant1x::pandas {
    // CROSS 金叉：前一根 S1 < S2 且当前 S1 > S2。
    inline std::vector<bool> cross(const Series &s1, const Series &s2) {
        std::size_t length = s1.size();
        if (s2.size() < length) {
            length = s2.size();
        }
        std::vector<bool> out(s1.size(), false);
        for (std::size_t i = 1; i < length; ++i) {
            out[i] = s1[i - 1] < s2[i - 1] && s1[i] > s2[i];
        }
        return out;
    }
}  // namespace quant1x::pandas

#endif  // PANDAS_FORMULA_CROSS_H
