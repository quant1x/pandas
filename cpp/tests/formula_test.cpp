// B2 首批公式的自包含测试：不依赖 gtest / catch2。
#include "formula/formula.h"

#include <algorithm>
#include <cmath>
#include <cstdio>
#include <limits>
#include <vector>

namespace {

using quant1x::pandas::Series;

int failures = 0;

bool close(double actual, double expected) {
    if (std::isnan(actual) && std::isnan(expected)) {
        return true;
    }
    if (actual == expected) {
        return true;
    }
    return std::abs(actual - expected) <= 1e-12 * std::max(1.0, std::abs(expected));
}

void expect_series(const char *label, const Series &got, const Series &want) {
    if (got.size() != want.size()) {
        std::printf("%s: size %zu != %zu\n", label, got.size(), want.size());
        ++failures;
        return;
    }
    for (std::size_t i = 0; i < want.size(); ++i) {
        if (!close(got[i], want[i])) {
            std::printf("%s: index %zu: %g != %g\n", label, i, got[i], want[i]);
            ++failures;
            return;
        }
    }
}

}  // namespace

int main() {
    const double nan = std::numeric_limits<double>::quiet_NaN();

    expect_series("abs", quant1x::pandas::abs({-1.5, 0.0, 2.5}), {1.5, 0.0, 2.5});

    const Series base = {1.0, 2.0, 3.0, 4.0, 5.0};
    expect_series("ref-1", quant1x::pandas::ref(base, 1), {nan, 1.0, 2.0, 3.0, 4.0});
    expect_series("ref-2", quant1x::pandas::ref(base, 2), {nan, nan, 1.0, 2.0, 3.0});
    expect_series("ref-neg", quant1x::pandas::ref(base, -1), {2.0, 3.0, 4.0, 5.0, nan});

    expect_series("ma-3", quant1x::pandas::ma({1.0, 2.0, 3.0, 4.0, 5.0}, 3),
                  {nan, nan, 2.0, 3.0, 4.0});
    expect_series("ma-empty", quant1x::pandas::ma({}, 3), {});

    expect_series("ema-4", quant1x::pandas::ema({1.0, 2.0, 3.0, 4.0}, 4),
                  {1.0, 1.4, 2.04, 2.824});

    const double alpha = 1.0 / 3.0;
    const double y1 = (1.0 - alpha) * 1.0 + alpha * 2.0;
    const double y2 = (1.0 - alpha) * y1 + alpha * 3.0;
    expect_series("sma-3-1", quant1x::pandas::sma({1.0, 2.0, 3.0}, 3, 1), {1.0, y1, y2});


    expect_series("hhv-2", quant1x::pandas::hhv({1.0, 3.0, 2.0, 5.0}, 2),
                  {nan, 3.0, 3.0, 5.0});
    expect_series("llv-2", quant1x::pandas::llv({1.0, 3.0, 2.0, 5.0}, 2),
                  {nan, 1.0, 2.0, 2.0});
    expect_series("sum-2", quant1x::pandas::sum({1.0, 2.0, 3.0, 4.0}, 2),
                  {nan, 3.0, 5.0, 7.0});
    expect_series("stddev-2", quant1x::pandas::stddev({1.0, 2.0, 3.0, 4.0}, 2),
                  {nan, 0.5, 0.5, 0.5});
    expect_series("count-2", quant1x::pandas::count({1.0, 0.0, 1.0, 1.0}, 2),
                  {0.0, 1.0, 1.0, 2.0});
    expect_series("wma-3", quant1x::pandas::wma({1.0, 2.0, 3.0}, 3),
                  {nan, nan, 14.0 / 6.0});
    expect_series("dma-2", quant1x::pandas::dma({1.0, 2.0, 3.0, 4.0}, {2.0}),
                  {1.0, 1.5, 2.25, 3.125});
    expect_series("dma-series", quant1x::pandas::dma({1.0, 2.0, 3.0}, {2.0, 4.0, 0.0}),
                  {1.0, 1.25, 3.0});

    const std::vector<bool> crossed = quant1x::pandas::cross({1.0, 3.0, 1.0, 3.0}, {2.0, 2.0, 2.0, 2.0});
    const std::vector<bool> want_cross = {false, true, false, true};
    if (crossed != want_cross) {
        std::printf("cross: mismatch\n");
        ++failures;
    }

    if (failures != 0) {
        std::printf("%d failure(s)\n", failures);
        return 1;
    }
    std::printf("all formula tests passed\n");
    return 0;
}
