const std = @import("std");
const support = @import("support.zig");

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
