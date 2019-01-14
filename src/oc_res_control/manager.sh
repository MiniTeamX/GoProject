#!/bin/bash

function start()
{
    ./xsafe.filter_service 1>/dev/null 2>&1 &
}

function stop()
{
    pid=`ps -ef | grep "xsafe.filter_service" | grep -v "grep" | awk '{print $2}'`
    kill -9 ${pid}
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
        sleep 1
        start
        ;;
esac

