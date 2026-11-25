#pragma once
#ifndef PANDAS_FORMULA_IF_H
#define PANDAS_FORMULA_IF_H 1

#include "support.h"

namespace quant1x::pandas {
    // IF / IFF 条件非零取 a，否则取 b；按最长对齐，缺位 NaN。
    inline Series if_select(const Series &condition, const Series &a, const Series &b) {
        std::size_t length = condition.size();
        if (a.size() > length) {
            length = a.size();
        }
        if (b.size() > length) {
            length = b.size();
        }
        Series out(length, nan_value());
        for (std::size_t i = 0; i < length; ++i) {
            const double flag = i < condition.size() ? condition[i] : 0.0;
            const double chosen = flag != 0.0 ? (i < a.size() ? a[i] : nan_value())
                                              : (i < b.size() ? b[i] : nan_value());
            out[i] = chosen;
        }
        return out;
    }
    // IFN 条件为零取 a，否则取 b。
    inline Series ifn(const Series &condition, const Series &a, const Series &b) {
        return if_select(condition, b, a);
    }
}  // namespace quant1x::pandas

#endif  // PANDAS_FORMULA_IF_H
