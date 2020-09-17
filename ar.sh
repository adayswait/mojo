#!/bin/bash
cd `dirname "$0"`
pidlist=$(pidof $(pwd)/mojo)
pidnum=${#pidlist[*]}
if [ ${pidnum} -gt 0 ]; then
	for svrpid in ${pidlist}
	do
		kill ${svrpid}
	done
fi
alias error=">&2 echo counter: "
# 以bash this.sh执行
cd `dirname "$0"`
nohup $(pwd)/mojo >/dev/null 2>&1 &
