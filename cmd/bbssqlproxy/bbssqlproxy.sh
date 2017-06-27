#!/bin/sh
#
# chkconfig: 2345 64 36
# description: bbssqlproxy startup scripts
#
bbsserver_root=/bbs/sbin
# each config file for one instance
configs="/bbs/etc/bbssqlproxy.ini"
pidfile="/bbs/etc/bbssqlproxy.pid"
nohupfile="/bbs/etc/nohup.bbssqlproxy.out"

ulimit -c unlimited
echo 1 > /proc/sys/fs/suid_dumpable
echo  "/bbs/share/%e-%p-%s-%t.core" >/proc/sys/kernel/core_pattern

ulimit -n 1024000

 startdebug() {
	if [ -f $pidfile ];then
		echo "bbssqlproxy is running"
		exit 7
	fi
    env GOTRACEBACK=crash nohup $bbsserver_root/bbssqlproxy -config $configs > $nohupfile &
    sleep 1
    test -f $pidfile && echo "Start bbssqlproxy success"
}

startX(){
	if [ -f $pidfile ];then
		echo "bbssqlproxy is running"
		exit 7
	fi
	nohup $bbsserver_root/bbssqlproxy -config $configs > $nohupfile &
        sleep 3
        if [ -f $pidfile ];then
            echo "Start bbssqlproxy success"
            rc=0
        else
            echo "bbssqlproxy is not start"
            rc=7
        fi
}
getpid(){
    pid=`ps -eo 'pid,args' | grep $bbsserver_root/bbssqlproxy | grep -v grep | awk '{print $1}'`
}
 
stopX(){
        getpid
        if [ -z "$pid" ];then
            echo "bbssqlproxy is not running"
            rc=0
        else
            kill $pid >/dev/null 2>&1
            sleep 3
            getpid
            if [ -n "$pid" ];then
                kill -9 $pid >/dev/null 2>&1
            fi
            echo "Stop bbssqlproxy success";
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
