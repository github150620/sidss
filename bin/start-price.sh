#!/bin/bash

cd `dirname $0`
SIDSS_HOME=$(pwd)/..

$SIDSS_HOME/bin/price >$SIDSS_HOME/log/price.log 2>&1 &
