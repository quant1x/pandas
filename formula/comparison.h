#pragma once
#ifndef PANDAS_FORMULA_COMPARISON_H
#define PANDAS_FORMULA_COMPARISON_H 1

#include "support.h"

namespace quant1x::pandas {
    // GT / GTE / LT / LTE / EQ / NEQ：按最长对齐，缺位 NaN，比较恒为 false。
    inline std::vector<bool> gt(const Series &x, const Series &y) {
        const std::size_t length = x.size() > y.size() ? x.size() : y.size();
        std::vector<bool> out(length, false);
        for (std::size_t i = 0; i < length; ++i) {
            out[i] = (i < x.size() ? x[i] : nan_value()) > (i < y.size() ? y[i] : nan_value());
        }
        return out;
    }
    inline std::vector<bool> gte(const Series &x, const Series &y) {
        const std::size_t length = x.size() > y.size() ? x.size() : y.size();
        std::vector<bool> out(length, false);
        for (std::size_t i = 0; i < length; ++i) {
            out[i] = (i < x.size() ? x[i] : nan_value()) >= (i < y.size() ? y[i] : nan_value());
        }
        return out;
    }
    inline std::vector<bool> lt(const Series &x, const Series &y) {
        const std::size_t length = x.size() > y.size() ? x.size() : y.size();
        std::vector<bool> out(length, false);
        for (std::size_t i = 0; i < length; ++i) {
            out[i] = (i < x.size() ? x[i] : nan_value()) < (i < y.size() ? y[i] : nan_value());
        }
        return out;
    }
    inline std::vector<bool> lte(const Series &x, const Series &y) {
        const std::size_t length = x.size() > y.size() ? x.size() : y.size();
        std::vector<bool> out(length, false);
        for (std::size_t i = 0; i < length; ++i) {
            out[i] = (i < x.size() ? x[i] : nan_value()) <= (i < y.size() ? y[i] : nan_value());
        }
        return out;
    }
    inline std::vector<bool> eq(const Series &x, const Series &y) {
        const std::size_t length = x.size() > y.size() ? x.size() : y.size();
        std::vector<bool> out(length, false);
        for (std::size_t i = 0; i < length; ++i) {
            out[i] = (i < x.size() ? x[i] : nan_value()) == (i < y.size() ? y[i] : nan_value());
        }
        return out;
    }
    inline std::vector<bool> neq(const Series &x, const Series &y) {
        const std::size_t length = x.size() > y.size() ? x.size() : y.size();
        std::vector<bool> out(length, false);
        for (std::size_t i = 0; i < length; ++i) {
            out[i] = (i < x.size() ? x[i] : nan_value()) != (i < y.size() ? y[i] : nan_value());
        }
        return out;
    }
    // AND / OR / NOT：逻辑族，缺位按 false。
    inline std::vector<bool> logic_and(const std::vector<bool> &x, const std::vector<bool> &y) {
        const std::size_t length = x.size() > y.size() ? x.size() : y.size();
        std::vector<bool> out(length, false);
        for (std::size_t i = 0; i < length; ++i) {
            out[i] = (i < x.size() && x[i]) && (i < y.size() && y[i]);
        }
        return out;
    }
    inline std::vector<bool> logic_or(const std::vector<bool> &x, const std::vector<bool> &y) {
        const std::size_t length = x.size() > y.size() ? x.size() : y.size();
        std::vector<bool> out(length, false);
        for (std::size_t i = 0; i < length; ++i) {
            out[i] = (i < x.size() && x[i]) || (i < y.size() && y[i]);
        }
        return out;
    }
    inline std::vector<bool> logic_not(const std::vector<bool> &x) {
        std::vector<bool> out(x.size(), false);
        for (std::size_t i = 0; i < x.size(); ++i) {
            out[i] = !x[i];
        }
        return out;
    }
}  // namespace quant1x::pandas

#endif  // PANDAS_FORMULA_COMPARISON_H
