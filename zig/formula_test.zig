//! 公式族 Zig 实现的单元用例与 CSV fixture 对拍。
//!
//! 在 `pandas/` 目录下执行 `zig test zig/formula_test.zig`；fixture 位于
//! `tests/formula/cases.csv`，与 Go / C++ / Rust 共用同一份期望值。

const std = @import("std");
const formula = @import("formula.zig");

fn nan() f64 {
    return std.math.nan(f64);
}

/// 读取 fixture：先按项目根目录解析，再按 `zig/` 目录解析。
fn readFixture(allocator: std.mem.Allocator) ![]u8 {
    const dir = std.Io.Dir.cwd();
    const limit = std.Io.Limit.limited(1 << 20);
    if (dir.readFileAlloc(std.testing.io, "tests/formula/cases.csv", allocator, limit)) |text| {
        return text;
    } else |_| {}
    return dir.readFileAlloc(std.testing.io, "../tests/formula/cases.csv", allocator, limit);
}

fn parseList(allocator: std.mem.Allocator, text: []const u8) ![]f64 {
    var list: std.ArrayList(f64) = .empty;
    errdefer list.deinit(allocator);
    var parts = std.mem.tokenizeScalar(u8, text, ';');
    while (parts.next()) |part| {
        const item = std.mem.trim(u8, part, " \t\r");
        const value = if (std.mem.eql(u8, item, "nan"))
            nan()
        else
            try std.fmt.parseFloat(f64, item);
        try list.append(allocator, value);
    }
    return list.toOwnedSlice(allocator);
}

fn parsePeriod(text: []const u8) !i64 {
    if (text.len == 0) return 0;
    return std.fmt.parseInt(i64, std.mem.trim(u8, text, " \t\r"), 10);
}

/// 逐值比对：NaN 对 NaN 视为相等，其余按相对 1e-12（下限 1）容差。
fn expectListEqual(name: []const u8, got: []const f64, want: []const f64) !void {
    if (got.len != want.len) {
        std.debug.print("{s}: 长度 {d} != {d}\n", .{ name, got.len, want.len });
        return error.LengthMismatch;
    }
    for (got, want, 0..) |actual, expected, index| {
        if (std.math.isNan(actual) and std.math.isNan(expected)) continue;
        const tolerance = 1e-12 * @max(@abs(expected), 1.0);
        if (@abs(actual - expected) <= tolerance) continue;
        std.debug.print("{s}: 第 {d} 位 {d} != {d}\n", .{ name, index, actual, expected });
        return error.ValueMismatch;
    }
}

test "abs 逐元素绝对值" {
    const allocator = std.testing.allocator;
    const x = [_]f64{ 1.0, -2.0, 3.0, -0.0, nan() };
    const got = try formula.abs(allocator, &x);
    defer allocator.free(got);
    try expectListEqual("abs", got, &[_]f64{ 1.0, 2.0, 3.0, 0.0, nan() });
}

test "ref 越界为 NaN" {
    const allocator = std.testing.allocator;
    const x = [_]f64{ 1.0, 2.0, 3.0, 4.0, 5.0 };
    const got = try formula.refValue(allocator, &x, 2);
    defer allocator.free(got);
    try expectListEqual("ref", got, &[_]f64{ nan(), nan(), 1.0, 2.0, 3.0 });

    const backward = try formula.refValue(allocator, &x, -1);
    defer allocator.free(backward);
    try expectListEqual("ref(-1)", backward, &[_]f64{ 2.0, 3.0, 4.0, 5.0, nan() });
}

test "ma 完整窗口才计算" {
    const allocator = std.testing.allocator;
    const x = [_]f64{ 1.0, 2.0, 3.0, 4.0, 5.0 };
    const got = try formula.ma(allocator, &x, 3);
    defer allocator.free(got);
    try expectListEqual("ma", got, &[_]f64{ nan(), nan(), 2.0, 3.0, 4.0 });

    const too_wide = try formula.ma(allocator, &x, 9);
    defer allocator.free(too_wide);
    try expectListEqual("ma(9)", too_wide, &[_]f64{ nan(), nan(), nan(), nan(), nan() });
}

test "ema 首值等于源序列首值" {
    const allocator = std.testing.allocator;
    const x = [_]f64{ 1.0, 2.0, 3.0, 4.0 };
    const got = try formula.ema(allocator, &x, 4);
    defer allocator.free(got);
    try expectListEqual("ema", got, &[_]f64{ 1.0, 1.4, 2.04, 2.824 });
}

test "sma 中国式平滑" {
    const allocator = std.testing.allocator;
    const x = [_]f64{ 1.0, 2.0, 3.0 };
    const got = try formula.sma(allocator, &x, 3, 1);
    defer allocator.free(got);
    try expectListEqual("sma", got, &[_]f64{ 1.0, 1.3333333333333335, 1.888888888888889 });
}

test "fixture cases.csv 逐行对拍" {
    const allocator = std.testing.allocator;
    const text = try readFixture(allocator);
    defer allocator.free(text);

    var lines = std.mem.splitScalar(u8, text, '\n');
    var checked: usize = 0;
    var header_seen = false;
    while (lines.next()) |raw| {
        const line = std.mem.trimEnd(u8, raw, "\r");
        if (line.len == 0) continue;
        if (!header_seen) {
            try std.testing.expectEqualStrings("formula,x,y,n,m,expected", line);
            header_seen = true;
            continue;
        }

        var fields: [6][]const u8 = undefined;
        var count: usize = 0;
        var cells = std.mem.splitScalar(u8, line, ',');
        while (cells.next()) |cell| {
            if (count == fields.len) return error.TooManyColumns;
            fields[count] = cell;
            count += 1;
        }
        if (count != fields.len) return error.TooFewColumns;

        const name = fields[0];
        const x = try parseList(allocator, fields[1]);
        defer allocator.free(x);
        const n = try parsePeriod(fields[3]);
        const m = try parsePeriod(fields[4]);
        const want = try parseList(allocator, fields[5]);
        defer allocator.free(want);

        const got = switch (name[0]) {
            'a' => try formula.abs(allocator, x),
            'r' => try formula.refValue(allocator, x, n),
            'm' => try formula.ma(allocator, x, n),
            'e' => try formula.ema(allocator, x, n),
            's' => try formula.sma(allocator, x, n, m),
            else => return error.UnknownFormula,
        };
        defer allocator.free(got);
        try expectListEqual(name, got, want);
        checked += 1;
    }
    try std.testing.expect(checked >= 5);
}
