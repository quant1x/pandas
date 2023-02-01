#!/usr/bin/env bash

set -e
git remote set-url origin https://github.com/quant1x/pandas.git
git push --all
git push --tags
git remote -vv