#pragma once
#ifndef PANDAS_FORMULA_H
#define PANDAS_FORMULA_H 1

// pandas 公式族的 C++ 对齐实现（B2）首批：ABS / REF / MA / EMA / SMA。
//
// 序列用 std::vector<double>，与 B1/B3 的最小参数集一致：
//   - REF 越界为 NaN；
//   - MA 只有完整窗口才计算，窗口不足为 NaN；
//   - EMA alpha=2/(N+1)、SMA alpha=M/N（M=0 视为 1），均按 EWM Adjust=false。

#include <algorithm>
#include <cmath>
#include <cstddef>
#include <cstdint>
#include <limits>
#include <vector>

namespace quant1x::pandas {

    using Series = std::vector<double>;

    inline double nan_value() { return std::numeric_limits<double>::quiet_NaN(); }

    // ABS 逐元素绝对值。
    inline Series abs(const Series &x) {
        Series out(x.size());
        for (std::size_t i = 0; i < x.size(); ++i) {
            out[i] = std::abs(x[i]);
        }
        return out;
    }

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

    // MA 简单移动平均：完整窗口才计算，窗口不足为 NaN。
    inline Series ma(const Series &x, std::size_t window) {
        Series out(x.size(), nan_value());
        if (window == 0 || window > x.size()) {
            return out;
        }
        double total = 0.0;
        for (std::size_t i = 0; i < x.size(); ++i) {
            total += x[i];
            if (i >= window) {
                total -= x[i - window];
            }
            if (i + 1 >= window) {
                out[i] = total / static_cast<double>(window);
            }
        }
        return out;
    }

    // EMA 指数移动平均：alpha = 2/(N+1)，首值等于源序列首值。
    inline Series ema(const Series &x, std::size_t span) {
        if (x.empty()) {
            return Series{};
        }
        const std::size_t period = span == 0 ? 1 : span;
        const double alpha = 2.0 / (static_cast<double>(period) + 1.0);
        Series out(x.size());
        out[0] = x[0];
        for (std::size_t i = 1; i < x.size(); ++i) {
            out[i] = (1.0 - alpha) * out[i - 1] + alpha * x[i];
        }
        return out;
    }

    // SMA 中国式平滑：alpha = M/N（M=0 视为 1）。
    inline Series sma(const Series &x, std::size_t n, int m) {
        if (x.empty()) {
            return Series{};
        }
        const std::size_t period = n == 0 ? 1 : n;
        const int weight = m == 0 ? 1 : m;
        const double alpha = static_cast<double>(weight) / static_cast<double>(period);
        Series out(x.size());
        out[0] = x[0];
        for (std::size_t i = 1; i < x.size(); ++i) {
            out[i] = (1.0 - alpha) * out[i - 1] + alpha * x[i];
        }
        return out;
    }


    // 滚动框架：完整窗口才计算，窗口不足为 NaN。
    template<typename F>
    inline Series rolling(const Series &x, std::size_t window, F apply) {
        Series out(x.size(), nan_value());
        if (window == 0 || window > x.size()) {
            return out;
        }
        for (std::size_t i = window - 1; i < x.size(); ++i) {
            out[i] = apply(x, i + 1 - window, i + 1);
        }
        return out;
    }

    // HHV 最近 N 个周期的最大值。
    inline Series hhv(const Series &x, std::size_t window) {
        return rolling(x, window, [](const Series &v, std::size_t begin, std::size_t end) {
            double value = v[begin];
            for (std::size_t i = begin + 1; i < end; ++i) {
                if (v[i] > value) {
                    value = v[i];
                }
            }
            return value;
        });
    }

    // LLV 最近 N 个周期的最小值。
    inline Series llv(const Series &x, std::size_t window) {
        return rolling(x, window, [](const Series &v, std::size_t begin, std::size_t end) {
            double value = v[begin];
            for (std::size_t i = begin + 1; i < end; ++i) {
                if (v[i] < value) {
                    value = v[i];
                }
            }
            return value;
        });
    }

    // HHVBARS 最近 N 周期内最高点到当前的周期数；并列取最近的一个。
    inline Series hhvbars(const Series &x, std::size_t window) {
        return rolling(x, window, [](const Series &v, std::size_t begin, std::size_t end) {
            std::size_t position = begin;
            for (std::size_t i = begin + 1; i < end; ++i) {
                if (v[i] >= v[position]) {
                    position = i;
                }
            }
            return static_cast<double>(end - 1 - position);
        });
    }

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

    // SUM 最近 N 个周期的和。
    inline Series sum(const Series &x, std::size_t window) {
        return rolling(x, window, [](const Series &v, std::size_t begin, std::size_t end) {
            double total = 0.0;
            for (std::size_t i = begin; i < end; ++i) {
                total += v[i];
            }
            return total;
        });
    }

    // STD（C++ 名 stddev）最近 N 个周期的总体标准差。
    inline Series stddev(const Series &x, std::size_t window) {
        return rolling(x, window, [](const Series &v, std::size_t begin, std::size_t end) {
            const double n = static_cast<double>(end - begin);
            double mean = 0.0;
            for (std::size_t i = begin; i < end; ++i) {
                mean += v[i];
            }
            mean /= n;
            double total = 0.0;
            for (std::size_t i = begin; i < end; ++i) {
                total += (v[i] - mean) * (v[i] - mean);
            }
            return std::sqrt(total / n);
        });
    }

    // COUNT 最近 N 个周期内非零条件的个数；窗口不足计 0。
    inline Series count(const Series &condition, std::size_t window) {
        Series out(condition.size(), 0.0);
        if (window == 0 || window > condition.size()) {
            return out;
        }
        for (std::size_t i = window - 1; i < condition.size(); ++i) {
            std::size_t total = 0;
            for (std::size_t j = i + 1 - window; j <= i; ++j) {
                if (condition[j] != 0.0) {
                    ++total;
                }
            }
            out[i] = static_cast<double>(total);
        }
        return out;
    }

    // WMA 加权移动平均：窗口内按 1..N 递增加权。
    inline Series wma(const Series &x, std::size_t window) {
        return rolling(x, window, [](const Series &v, std::size_t begin, std::size_t end) {
            double total = 0.0;
            double weights = 0.0;
            for (std::size_t i = begin; i < end; ++i) {
                const double weight = static_cast<double>(i - begin + 1);
                total += weight * v[i];
                weights += weight;
            }
            return total / weights;
        });
    }

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


    // MAX 两条序列逐元素取大，按最长对齐，缺位为 NaN。
    inline Series max_series(const Series &x, const Series &y) {
        const std::size_t length = x.size() > y.size() ? x.size() : y.size();
        Series out(length, nan_value());
        for (std::size_t i = 0; i < length; ++i) {
            const double a = i < x.size() ? x[i] : nan_value();
            const double b = i < y.size() ? y[i] : nan_value();
            out[i] = b > a ? b : a;
        }
        return out;
    }

    // MIN 两条序列逐元素取小。
    inline Series min_series(const Series &x, const Series &y) {
        const std::size_t length = x.size() > y.size() ? x.size() : y.size();
        Series out(length, nan_value());
        for (std::size_t i = 0; i < length; ++i) {
            const double a = i < x.size() ? x[i] : nan_value();
            const double b = i < y.size() ? y[i] : nan_value();
            out[i] = b < a ? b : a;
        }
        return out;
    }

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

    // BARSLAST 上一次条件成立到当前的周期数；从未成立为 -1。
    inline Series barslast(const Series &condition) {
        Series out(condition.size(), -1.0);
        std::int64_t last_true = -1;
        for (std::size_t i = 0; i < condition.size(); ++i) {
            if (condition[i] != 0.0) {
                last_true = static_cast<std::int64_t>(i);
                out[i] = 0.0;
            } else if (last_true >= 0) {
                out[i] = static_cast<double>(static_cast<std::int64_t>(i) - last_true);
            }
        }
        return out;
    }

    // BARSLASTCOUNT 连续满足条件的周期数。
    inline Series barslastcount(const Series &condition) {
        Series out(condition.size(), 0.0);
        double streak = 0.0;
        for (std::size_t i = 0; i < condition.size(); ++i) {
            if (condition[i] != 0.0) {
                streak += 1.0;
                out[i] = streak;
            } else {
                streak = 0.0;
            }
        }
        return out;
    }

    // BARSSINCEN 窗口内第一次条件成立到当前的周期数；没有成立为 0。
    inline Series barssincen(const Series &condition, std::size_t window) {
        Series out(condition.size(), 0.0);
        if (window == 0 || window > condition.size()) {
            return out;
        }
        for (std::size_t i = window - 1; i < condition.size(); ++i) {
            const std::size_t begin = i + 1 - window;
            for (std::size_t j = begin; j <= i; ++j) {
                if (condition[j] != 0.0) {
                    out[i] = static_cast<double>(i - j);
                    break;
                }
            }
        }
        return out;
    }

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

    // SQRT 逐元素平方根。
    inline Series sqrt(const Series &x) {
        Series out(x.size());
        for (std::size_t i = 0; i < x.size(); ++i) {
            out[i] = std::sqrt(x[i]);
        }
        return out;
    }


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

    // LAST 持续存在：A+1 窗口内排除最近 B 根后全部非零；窗口不足 NaN。
    inline Series last(const Series &condition, std::size_t a, std::size_t b) {
        return rolling(condition, a + 1, [b](const Series &v, std::size_t begin, std::size_t end) {
            if (b >= end - begin) {
                return 1.0;
            }
            const std::size_t stop = end - b;
            for (std::size_t i = begin; i < stop; ++i) {
                if (v[i] == 0.0) {
                    return 0.0;
                }
            }
            return 1.0;
        });
    }

    // BARSLASTS 倒数第 N 次成立到当前的周期数（Go 递推口径）。
    inline Series barslasts(const Series &condition, std::int64_t n) {
        Series current(condition.size(), 0.0);
        double previous = 0.0;
        for (std::size_t i = 0; i < condition.size(); ++i) {
            previous = condition[i] != 0.0 ? 0.0 : previous + 1.0;
            current[i] = previous;
        }
        for (std::int64_t iteration = 1; iteration < n; ++iteration) {
            Series next(current.size(), nan_value());
            for (std::size_t i = 0; i < current.size(); ++i) {
                if (!std::isfinite(current[i])) {
                    continue;
                }
                const double step = current[i] + 1.0;
                const double source = static_cast<double>(i) - step;
                if (source < 0.0) {
                    continue;
                }
                const std::size_t index = static_cast<std::size_t>(source);
                if (index >= current.size() || !std::isfinite(current[index])) {
                    continue;
                }
                next[i] = current[index] + step;
            }
            current = next;
        }
        return current;
    }

}  // namespace quant1x::pandas

#endif  // PANDAS_FORMULA_H
