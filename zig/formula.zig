//! pandas 公式族的 Zig 扩展实现（B1 首批）。
//!
//! 覆盖 ABS / REF / MA / EMA / SMA；序列用 `[]const f64` + `std.mem.Allocator`，
//! 与 Go / C++ / Rust 端的最小参数集一致。EMA / SMA 按 Go 的 EWM `Adjust=false` 口径：
//! `y0=x0`、`y_t=(1-alpha)*y_{t-1}+alpha*x_t`，NaN 位置沿用前值。

const std = @import("std");

fn nan() f64 {
    return std.math.nan(f64);
}

/// ABS 逐元素绝对值。
pub fn abs(allocator: std.mem.Allocator, x: []const f64) ![]f64 {
    const out = try allocator.alloc(f64, x.len);
    for (x, 0..) |value, index| {
        out[index] = @abs(value);
    }
    return out;
}

/// REF 引用前 N 个周期的值；越界为 NaN。
pub fn refValue(allocator: std.mem.Allocator, x: []const f64, n: i64) ![]f64 {
    const out = try allocator.alloc(f64, x.len);
    if (n == 0) {
        @memcpy(out, x);
        return out;
    }
    @memset(out, nan());
    const distance: i128 = if (n < 0) -@as(i128, n) else @as(i128, n);
    const width: usize = @intCast(@min(distance, @as(i128, @intCast(x.len))));
    if (n > 0) {
        @memcpy(out[width..], x[0 .. x.len - width]);
    } else {
        @memcpy(out[0 .. x.len - width], x[width..]);
    }
    return out;
}

/// MA 简单移动平均：完整窗口才计算，窗口不足为 NaN。
pub fn ma(allocator: std.mem.Allocator, x: []const f64, window: i64) ![]f64 {
    const out = try allocator.alloc(f64, x.len);
    @memset(out, nan());
    if (window <= 0) return out;
    const size: usize = @intCast(window);
    if (size > x.len) return out;
    var index: usize = size - 1;
    while (index < x.len) : (index += 1) {
        const start = index + 1 - size;
        var total: f64 = 0;
        for (x[start .. index + 1]) |value| {
            total += value;
        }
        out[index] = total / @as(f64, @floatFromInt(size));
    }
    return out;
}

/// EMA 指数移动平均：alpha = 2/(N+1)，首值等于源序列首值。
pub fn ema(allocator: std.mem.Allocator, x: []const f64, span: i64) ![]f64 {
    const period: f64 = @floatFromInt(@max(span, 1));
    return ewm(allocator, x, 2.0 / (period + 1.0));
}

/// SMA 中国式平滑：alpha = M/N（M=0 视为 1）。
pub fn sma(allocator: std.mem.Allocator, x: []const f64, n: i64, m: i64) ![]f64 {
    const weight: f64 = @floatFromInt(if (m == 0) @as(i64, 1) else m);
    const period: f64 = @floatFromInt(@max(n, 1));
    return ewm(allocator, x, weight / period);
}

/// HHV 最近 N 个周期的最大值。
pub fn hhv(allocator: std.mem.Allocator, x: []const f64, window: i64) ![]f64 {
    return rolling(maxOf, allocator, x, window);
}

/// LLV 最近 N 个周期的最小值。
pub fn llv(allocator: std.mem.Allocator, x: []const f64, window: i64) ![]f64 {
    return rolling(minOf, allocator, x, window);
}

/// SUM 最近 N 个周期的和。
pub fn sum(allocator: std.mem.Allocator, x: []const f64, window: i64) ![]f64 {
    return rolling(totalOf, allocator, x, window);
}

/// STD 最近 N 个周期的总体标准差（符号名 `stddev` 避开 `std`）。
pub fn stddev(allocator: std.mem.Allocator, x: []const f64, window: i64) ![]f64 {
    return rolling(deviationOf, allocator, x, window);
}

/// WMA 加权移动平均：窗口内按 1..N 递增加权，分母 N(N+1)/2。
pub fn wma(allocator: std.mem.Allocator, x: []const f64, window: i64) ![]f64 {
    return rolling(weightedMeanOf, allocator, x, window);
}

/// SQRT 逐元素平方根。
pub fn sqrt(allocator: std.mem.Allocator, x: []const f64) ![]f64 {
    const out = try allocator.alloc(f64, x.len);
    for (x, 0..) |value, index| {
        out[index] = @sqrt(value);
    }
    return out;
}

/// COUNT 最近 N 个周期内非零条件的个数；窗口不足计 0。
pub fn count(allocator: std.mem.Allocator, condition: []const f64, window: i64) ![]f64 {
    const out = try allocator.alloc(f64, condition.len);
    @memset(out, 0);
    if (window <= 0) return out;
    const size: usize = @intCast(window);
    if (size > condition.len) return out;
    var index: usize = size - 1;
    while (index < condition.len) : (index += 1) {
        var total: usize = 0;
        for (condition[index + 1 - size .. index + 1]) |value| {
            if (value != 0) total += 1;
        }
        out[index] = @floatFromInt(total);
    }
    return out;
}

/// HHVBARS 最近 N 周期内最高点到当前的周期数；并列取最近的一个。
pub fn hhvbars(allocator: std.mem.Allocator, x: []const f64, window: i64) ![]f64 {
    return rolling(barsSinceMax, allocator, x, window);
}

/// LLVBARS 最近 N 周期内最低点到当前的周期数；并列取最近的一个。
pub fn llvbars(allocator: std.mem.Allocator, x: []const f64, window: i64) ![]f64 {
    return rolling(barsSinceMin, allocator, x, window);
}

/// CROSS 金叉：前一根 S1 < S2 且当前 S1 > S2。
pub fn cross(allocator: std.mem.Allocator, s1: []const f64, s2: []const f64) ![]bool {
    const out = try allocator.alloc(bool, s1.len);
    @memset(out, false);
    const length = @min(s1.len, s2.len);
    var index: usize = 1;
    while (index < length) : (index += 1) {
        out[index] = s1[index - 1] < s2[index - 1] and s1[index] > s2[index];
    }
    return out;
}

/// DMA 动态移动平均：alpha = 1/A，A 可常量或逐位置；非有限时按 1。
pub fn dma(allocator: std.mem.Allocator, x: []const f64, a: []const f64) ![]f64 {
    const out = try allocator.alloc(f64, x.len);
    if (x.len == 0) return out;
    out[0] = x[0];
    var index: usize = 1;
    while (index < x.len) : (index += 1) {
        const period = if (a.len == 1) a[0] else a[index];
        var alpha = 1.0 / period;
        if (!std.math.isFinite(alpha)) alpha = 1.0;
        out[index] = alpha * x[index] + (1 - alpha) * out[index - 1];
    }
    return out;
}

/// FILTER 条件成立后 N 位清零；扫描已处理的序列。
pub fn filter(allocator: std.mem.Allocator, condition: []const f64, window: i64) ![]f64 {
    const out = try allocator.alloc(f64, condition.len);
    @memcpy(out, condition);
    if (window <= 0) return out;
    const width: usize = @intCast(window);
    var index: usize = 0;
    while (index < out.len) : (index += 1) {
        if (out[index] == 0) continue;
        const end = @min(index + 1 + width, out.len);
        @memset(out[index + 1 .. end], 0);
    }
    return out;
}

/// MAX 两条序列逐元素取大；要求等长。
pub fn maxSeries(allocator: std.mem.Allocator, x: []const f64, y: []const f64) ![]f64 {
    if (x.len != y.len) return error.LengthMismatch;
    const out = try allocator.alloc(f64, x.len);
    for (x, y, 0..) |left, right, index| {
        out[index] = @max(left, right);
    }
    return out;
}

/// MIN 两条序列逐元素取小；要求等长。
pub fn minSeries(allocator: std.mem.Allocator, x: []const f64, y: []const f64) ![]f64 {
    if (x.len != y.len) return error.LengthMismatch;
    const out = try allocator.alloc(f64, x.len);
    for (x, y, 0..) |left, right, index| {
        out[index] = @min(left, right);
    }
    return out;
}

/// GT 逐元素大于。
pub fn gt(allocator: std.mem.Allocator, x: []const f64, y: []const f64) ![]bool {
    return compareWith(greaterThan, allocator, x, y);
}

/// GTE 逐元素大于等于。
pub fn gte(allocator: std.mem.Allocator, x: []const f64, y: []const f64) ![]bool {
    return compareWith(greaterOrEqual, allocator, x, y);
}

/// LT 逐元素小于。
pub fn lt(allocator: std.mem.Allocator, x: []const f64, y: []const f64) ![]bool {
    return compareWith(lessThan, allocator, x, y);
}

/// LTE 逐元素小于等于。
pub fn lte(allocator: std.mem.Allocator, x: []const f64, y: []const f64) ![]bool {
    return compareWith(lessOrEqual, allocator, x, y);
}

/// EQ 逐元素相等；NaN 与自身不相等。
pub fn eq(allocator: std.mem.Allocator, x: []const f64, y: []const f64) ![]bool {
    return compareWith(equalTo, allocator, x, y);
}

/// NEQ 逐元素不等；NaN 与自身视为不等。
pub fn neq(allocator: std.mem.Allocator, x: []const f64, y: []const f64) ![]bool {
    return compareWith(notEqualTo, allocator, x, y);
}

/// AND 逻辑与（`and` 是关键字，落在 `logicAnd`）。
pub fn logicAnd(allocator: std.mem.Allocator, x: []const bool, y: []const bool) ![]bool {
    if (x.len != y.len) return error.LengthMismatch;
    const out = try allocator.alloc(bool, x.len);
    for (x, y, 0..) |left, right, index| {
        out[index] = left and right;
    }
    return out;
}

/// OR 逻辑或（`or` 是关键字，落在 `logicOr`）。
pub fn logicOr(allocator: std.mem.Allocator, x: []const bool, y: []const bool) ![]bool {
    if (x.len != y.len) return error.LengthMismatch;
    const out = try allocator.alloc(bool, x.len);
    for (x, y, 0..) |left, right, index| {
        out[index] = left or right;
    }
    return out;
}

/// NOT 逻辑非。
pub fn logicNot(allocator: std.mem.Allocator, x: []const bool) ![]bool {
    const out = try allocator.alloc(bool, x.len);
    for (x, 0..) |value, index| {
        out[index] = !value;
    }
    return out;
}

fn compareWith(
    comptime op: fn (f64, f64) bool,
    allocator: std.mem.Allocator,
    x: []const f64,
    y: []const f64,
) ![]bool {
    if (x.len != y.len) return error.LengthMismatch;
    const out = try allocator.alloc(bool, x.len);
    for (x, y, 0..) |left, right, index| {
        out[index] = op(left, right);
    }
    return out;
}

fn greaterThan(a: f64, b: f64) bool {
    return a > b;
}

fn greaterOrEqual(a: f64, b: f64) bool {
    return a >= b;
}

fn lessThan(a: f64, b: f64) bool {
    return a < b;
}

fn lessOrEqual(a: f64, b: f64) bool {
    return a <= b;
}

fn equalTo(a: f64, b: f64) bool {
    return a == b;
}

fn notEqualTo(a: f64, b: f64) bool {
    return a != b;
}

/// IF 条件选择：条件非零取 A，否则取 B；三条序列等长。
pub fn ifSelect(
    allocator: std.mem.Allocator,
    condition: []const f64,
    a: []const f64,
    b: []const f64,
) ![]f64 {
    if (condition.len != a.len or condition.len != b.len) return error.LengthMismatch;
    const out = try allocator.alloc(f64, condition.len);
    for (condition, a, b, 0..) |flag, left, right, index| {
        out[index] = if (flag != 0) left else right;
    }
    return out;
}

/// IFN 反向条件选择：条件非零取 B，否则取 A。
pub fn ifn(
    allocator: std.mem.Allocator,
    condition: []const f64,
    a: []const f64,
    b: []const f64,
) ![]f64 {
    return ifSelect(allocator, condition, b, a);
}

/// BARSLAST 上一次条件成立到当前的周期数；从未成立为 -1。
pub fn barslast(allocator: std.mem.Allocator, condition: []const f64) ![]f64 {
    const out = try allocator.alloc(f64, condition.len);
    var last_true: ?usize = null;
    for (condition, 0..) |value, index| {
        if (value != 0) {
            last_true = index;
            out[index] = 0;
        } else if (last_true) |position| {
            out[index] = @floatFromInt(index - position);
        } else {
            out[index] = -1;
        }
    }
    return out;
}

/// BARSLASTCOUNT 连续成立的周期数；遇到不成立归零。
pub fn barslastcount(allocator: std.mem.Allocator, condition: []const f64) ![]f64 {
    const out = try allocator.alloc(f64, condition.len);
    var streak: f64 = 0;
    for (condition, 0..) |value, index| {
        streak = if (value != 0) streak + 1 else 0;
        out[index] = streak;
    }
    return out;
}

/// BARSSINCEN N 周期内第一次成立到当前的周期数；窗口不足为 0。
pub fn barssincen(allocator: std.mem.Allocator, condition: []const f64, window: i64) ![]f64 {
    const out = try allocator.alloc(f64, condition.len);
    @memset(out, 0);
    if (window <= 0) return out;
    const size: usize = @intCast(window);
    if (size > condition.len) return out;
    var index: usize = size - 1;
    while (index < condition.len) : (index += 1) {
        const begin = index + 1 - size;
        for (condition[begin .. index + 1], begin..) |value, position| {
            if (value != 0) {
                out[index] = @floatFromInt(index - position);
                break;
            }
        }
    }
    return out;
}

/// BARSLASTS 倒数第 N 次成立到当前的周期数；越界为 NaN。
pub fn barslasts(allocator: std.mem.Allocator, condition: []const f64, n: i64) ![]f64 {
    var current = try barslastZero(allocator, condition);
    var iteration: i64 = 1;
    while (iteration < n) : (iteration += 1) {
        const following = try allocator.alloc(f64, current.len);
        @memset(following, nan());
        for (current, 0..) |value, index| {
            if (std.math.isNan(value)) continue;
            const step = value + 1;
            const source = @as(f64, @floatFromInt(index)) - step;
            if (source < 0) continue;
            const source_index: usize = @intFromFloat(source);
            if (source_index >= current.len) continue;
            if (std.math.isNan(current[source_index])) continue;
            following[index] = current[source_index] + step;
        }
        allocator.free(current);
        current = following;
    }
    return current;
}

/// LAST 持续存在：A+1 窗口内排除最近 B 根后全部非零；窗口不足为 NaN。
pub fn last(allocator: std.mem.Allocator, condition: []const f64, a: i64, b: i64) ![]f64 {
    const out = try allocator.alloc(f64, condition.len);
    @memset(out, nan());
    if (a < 0) return out;
    const size: usize = @intCast(a + 1);
    if (size == 0 or size > condition.len) return out;
    const skip: usize = if (b < 0) 0 else @intCast(b);
    var index: usize = size - 1;
    while (index < condition.len) : (index += 1) {
        if (skip >= size) {
            out[index] = 1;
            continue;
        }
        const begin = index + 1 - size;
        const stop = index + 1 - skip;
        var satisfied = true;
        for (condition[begin..stop]) |value| {
            if (value == 0) {
                satisfied = false;
                break;
            }
        }
        out[index] = if (satisfied) 1 else 0;
    }
    return out;
}

fn barslastZero(allocator: std.mem.Allocator, condition: []const f64) ![]f64 {
    const out = try allocator.alloc(f64, condition.len);
    var previous: f64 = 0;
    for (condition, 0..) |value, index| {
        previous = if (value != 0) 0 else previous + 1;
        out[index] = previous;
    }
    return out;
}

/// 逐窗口应用：完整窗口才计算，窗口不足为 NaN。
fn rolling(
    comptime apply: fn ([]const f64) f64,
    allocator: std.mem.Allocator,
    x: []const f64,
    window: i64,
) ![]f64 {
    const out = try allocator.alloc(f64, x.len);
    @memset(out, nan());
    if (window <= 0) return out;
    const size: usize = @intCast(window);
    if (size > x.len) return out;
    var index: usize = size - 1;
    while (index < x.len) : (index += 1) {
        out[index] = apply(x[index + 1 - size .. index + 1]);
    }
    return out;
}

fn maxOf(block: []const f64) f64 {
    var best = block[0];
    for (block[1..]) |value| {
        if (value > best) best = value;
    }
    return best;
}

fn minOf(block: []const f64) f64 {
    var best = block[0];
    for (block[1..]) |value| {
        if (value < best) best = value;
    }
    return best;
}

fn totalOf(block: []const f64) f64 {
    var total: f64 = 0;
    for (block) |value| total += value;
    return total;
}

fn deviationOf(block: []const f64) f64 {
    const size: f64 = @floatFromInt(block.len);
    const mean = totalOf(block) / size;
    var squares: f64 = 0;
    for (block) |value| {
        const delta = value - mean;
        squares += delta * delta;
    }
    return @sqrt(squares / size);
}

fn weightedMeanOf(block: []const f64) f64 {
    var total: f64 = 0;
    var weights: f64 = 0;
    for (block, 0..) |value, index| {
        const weight: f64 = @floatFromInt(index + 1);
        total += weight * value;
        weights += weight;
    }
    return total / weights;
}

fn barsSinceMax(block: []const f64) f64 {
    var position: usize = 0;
    var best = block[0];
    for (block, 0..) |value, index| {
        if (value >= best) {
            best = value;
            position = index;
        }
    }
    return @floatFromInt(block.len - 1 - position);
}

fn barsSinceMin(block: []const f64) f64 {
    var position: usize = 0;
    var best = block[0];
    for (block, 0..) |value, index| {
        if (value <= best) {
            best = value;
            position = index;
        }
    }
    return @floatFromInt(block.len - 1 - position);
}

/// EWM `Adjust=false` 递归：首值 NaN 视为 0，NaN 位置沿用前值。
fn ewm(allocator: std.mem.Allocator, x: []const f64, alpha: f64) ![]f64 {
    const out = try allocator.alloc(f64, x.len);
    if (x.len == 0) return out;
    var previous = x[0];
    if (std.math.isNan(previous)) previous = 0;
    out[0] = previous;
    const beta = 1 - alpha;
    var index: usize = 1;
    while (index < x.len) : (index += 1) {
        const value = x[index];
        if (std.math.isNan(value)) {
            out[index] = previous;
            continue;
        }
        previous = beta * previous + alpha * value;
        if (std.math.isNan(previous)) previous = out[index - 1];
        out[index] = previous;
    }
    return out;
}
