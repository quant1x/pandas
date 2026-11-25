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

/// EWM `Adjust=false` 递归：首值 NaN 视为 0，NaN 位置沿用前值。
fn ewm(allocator: std.mem.Allocator, x: []const f64, alpha: f64) ![]f64 {
    const out = try allocator.alloc(f64, x.len);
    if (x.len == 0) return out;
    var last = x[0];
    if (std.math.isNan(last)) last = 0;
    out[0] = last;
    const beta = 1 - alpha;
    var index: usize = 1;
    while (index < x.len) : (index += 1) {
        const value = x[index];
        if (std.math.isNan(value)) {
            out[index] = last;
            continue;
        }
        last = beta * last + alpha * value;
        if (std.math.isNan(last)) last = out[index - 1];
        out[index] = last;
    }
    return out;
}
