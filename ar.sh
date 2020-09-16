#!/bin/bash
alias error=">&2 echo counter: "
# 以bash this.sh执行
cd `dirname "$0"`
nohup $(pwd)/mojo >/dev/null 2>&1 &
echo mojoarok
