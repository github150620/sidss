#!/bin/bash

cd `dirname $0`
SIDSS_HOME=$(pwd)/..

$SIDSS_HOME/bin/kline >$SIDSS_HOME/log/kline.log 2>&1 &
