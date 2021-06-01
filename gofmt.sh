#!/bin/bash
# 格式化用户目录下全部go文件
find /Users/zen -name "*.go" -exec gofmt -w {} \;