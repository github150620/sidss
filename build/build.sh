#!/bin/bash

cd `dirname $0`
SIDSS_SRC=$(pwd)/../src
SIDSS_BIN=$(pwd)/../bin

go build -i -o $SIDSS_BIN/alarm $SIDSS_SRC/cmd/alarm/main.go
go build -i -o $SIDSS_BIN/kline $SIDSS_SRC/cmd/kline/main.go
go build -i -o $SIDSS_BIN/ma    $SIDSS_SRC/cmd/ma/main.go
go build -i -o $SIDSS_BIN/price $SIDSS_SRC/cmd/price/main.go

