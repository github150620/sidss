#!/bin/bash

cd `dirname $0`
SIDSS_HOME=$(pwd)/..

$SIDSS_HOME/bin/ma >$SIDSS_HOME/log/ma.log 2>&1 &
