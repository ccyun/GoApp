#!/bin/sh
#
# chkconfig: 2345 64 36
# description: bbstaskserver startup scripts
#
bbsserver_root=/bbs/sbin
# each config file for one instance
configs="/bbs/etc/bbstaskserver.ini"
pidfile="/bbs/etc/bbstaskserver.pid"
nohupfile="/bbs/etc/nohup.bbstaskserver.out"

ulimit -c unlimited
echo 1 > /proc/sys/fs/suid_dumpable
echo  "/bbs/share/%e-%p-%s-%t.core" >/proc/sys/kernel/core_pattern

ulimit -n 1024000

 startdebug() {
	if [ -f $pidfile ];then
		echo "bbstaskserver is running"
		exit 7
	fi
    env GOTRACEBACK=crash nohup $bbsserver_root/bbstaskserver -c $configs > $nohupfile &
    sleep 1
    getpid
    echo $pid>>$pidfile
    test -f $pidfile && echo "Start bbstaskserver success"
}

startX(){
	if [ -f $pidfile ];then
		echo "bbstaskserver is running"
		exit 7
	fi
	nohup $bbsserver_root/bbstaskserver -c $configs > $nohupfile &
        sleep 3
        getpid
        echo $pid>>$pidfile
        if [ -f $pidfile ];then
            echo "Start bbstaskserver success"
            rc=0
        else
            echo "bbstaskserver is not start"
            rc=7
        fi
}
getpid(){
    pid=`ps -eo 'pid,args' | grep $bbsserver_root/bbstaskserver | grep -v grep | awk '{print $1}'`
}
 
stopX(){
        getpid
        if [ -z "$pid" ];then
            echo "bbstaskserver is not running"
            rc=0
        else
            kill $pid >/dev/null 2>&1
            sleep 3
            getpid
            if [ -n "$pid" ];then
                kill -9 $pid >/dev/null 2>&1
            fi
            echo "Stop bbstaskserver success";
        fi
	test -f $pidfile && rm $pidfile
}
 
# See how we were called.
case "$1" in
    start)
        startX
        ;;
    startdebug)
        startdebug
        ;;
    stop)
        stopX
        ;;
    restart)
        stopX
        startX
        ;;
    *)
        echo $"Usage: $0 {start|stop|restart}"
        ;;
esac
exit 0 
