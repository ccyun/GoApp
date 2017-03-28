#!/bin/sh
#
# chkconfig: 2345 64 36
# description: bbsapp quere script  startup scripts
bbstaskscript_root="/home/ubuntu/cmd/TaskScript"
configs="/home/ubuntu/cmd/TaskScript/config.ini"
nohupfile="/home/ubuntu/cmd/TaskScript/nohup.out"

ulimit -c unlimited
echo 1 > /proc/sys/fs/suid_dumpable
echo  "/uc/share/%e-%p-%s-%t.core" >/proc/sys/kernel/core_pattern

ulimit -n 1024000 
 
start() {
	nohup $bbstaskscript_root/TaskScript -c $configs >> $nohupfile &
	sleep 1
	test -f|pgrep TaskScript>/dev/null && echo "Start bbsapp taskscript success"
}
 
stop() {
	pgrep TaskScript|xargs kill -9 >/dev/null 2>&1 &
}
 
case "$1" in
    start)
        start
        ;;
    stop)
        stop
        ;;
    restart)
        stop
        start
        ;;
    *)
        echo $"Usage: $0 {start|stop|restart}"
        ;;
esac
exit 0 

