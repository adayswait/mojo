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
echo mojobrok
