#!/bin/bash

cd `dirname $0`
SIDSS_HOME=$(pwd)/..

$SIDSS_HOME/bin/alarm >$SIDSS_HOME/log/alarm.log 2>&1 &
