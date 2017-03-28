#!/bin/sh
#
# chkconfig: 2345 64 36
# description: bbsapp quere script  startup scripts
root="/home/ubuntu/cmd/TaskScript"
configs=$root"/conf.ini"
nohupfile=$root"/nohup.out"

ulimit -c unlimited
echo 1 > /proc/sys/fs/suid_dumpable
echo  "/uc/share/%e-%p-%s-%t.core" >/proc/sys/kernel/core_pattern

ulimit -n 1024000 
 
start() {
    if test $(pgrep -f TaskScript|wc -l) -eq 0
    then
        echo "taskscript is running..."
    else
        $root/TaskScript -c $configs >> $nohupfile &
        return
    fi
	sleep 2
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

