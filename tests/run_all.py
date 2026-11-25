#!/usr/bin/env python3
"""pandas 仓的跨语言验收：生成器幂等 + Go / C++ / Rust / Zig 四端。

用法（仓库根）：

    python tests/run_all.py            # 全部
    python tests/run_all.py --only go  # 只跑名字含 go 的步骤

退出码：全部通过为 0，任一步失败为 1。
"""

from __future__ import annotations

import argparse
import hashlib
import os
import subprocess
import sys
import time
from pathlib import Path

ROOT = Path(__file__).resolve().parent.parent
PYTHON = sys.executable

# 生成器与其产物：跑完必须逐字节一致（幂等）。
GENERATORS = [
    ("tests/formula/gen_fixtures.py", ["tests/formula/cases.csv"]),
    ("tests/features/gen_features.py",
     ["tests/features/features.csv", "tests/features/model.csv", "tests/features/model_mlp.csv"]),
    ("scripts/factor_report.py", ["tests/features/report.csv"]),
]


def digest(paths: list[str]) -> dict[str, str]:
    out = {}
    for relative in paths:
        data = (ROOT / relative).read_bytes()
        out[relative] = hashlib.sha256(data).hexdigest()
    return out


def run(command: list[str]) -> None:
    subprocess.run(command, cwd=ROOT, check=True)


def check_generators() -> None:
    before = {script: digest(files) for script, files in GENERATORS}
    for script, _ in GENERATORS:
        run([PYTHON, script])
    after = {script: digest(files) for script, files in GENERATORS}
    for script, files in GENERATORS:
        if before[script] != after[script]:
            raise RuntimeError(f"{script} 生成的产物不幂等：{files}")


def onnx_step() -> str:
    """ONNX 推理：需要 ORT_DYLIB_PATH；未设置时返回跳过原因。"""
    dylib = os.environ.get("ORT_DYLIB_PATH", "")
    if not dylib:
        return "未设置 ORT_DYLIB_PATH（onnxruntime 动态库）"
    run(["cargo", "test", "--manifest-path", "pandas-rust/Cargo.toml", "--test", "onnx"])
    return ""


STEPS = [
    ("fixtures 幂等", check_generators),
    ("go 公式 fixture", lambda: run(["go", "test", "./formula/"])),
    ("go 特征 fixture", lambda: run(["go", "test", "./tests/features/"])),
    ("cpp ctest", lambda: run(["ctest", "--test-dir", "cpp/build", "-C", "Debug",
                               "--output-on-failure"])),
    ("rust 用例", lambda: run(["cargo", "test", "--manifest-path",
                               "pandas-rust/Cargo.toml"])),
    ("zig 用例", lambda: run(["zig", "test", "zig/formula_test.zig"])),
    ("rust onnx", onnx_step),
]


def main() -> int:
    parser = argparse.ArgumentParser(description="跨语言验收")
    parser.add_argument("--only", default="", help="只跑名字包含该子串的步骤")
    args = parser.parse_args()

    failures = []
    for name, action in STEPS:
        if args.only and args.only not in name:
            continue
        started = time.time()
        try:
            reason = action()
        except (subprocess.CalledProcessError, RuntimeError) as error:
            failures.append(name)
            print(f"FAIL  {name}  ({time.time() - started:.1f}s)  {error}")
        else:
            if reason:
                print(f"SKIP  {name}  {reason}")
            else:
                print(f"PASS  {name}  ({time.time() - started:.1f}s)")

    if failures:
        print(f"\n{len(failures)} 步失败：{', '.join(failures)}")
        return 1
    print("\n全部通过")
    return 0


if __name__ == "__main__":
    sys.exit(main())
