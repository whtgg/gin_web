#!/usr/bin/env bash
APP_DIR=$(
  cd $(dirname $0)
  pwd
)

APP_NAME="${APP_DIR##*/}-server"

TARGET_DIR=~/Downloads

# mac 环境编译为linux环境
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags=jsoniter -ldflags "-s -w" -o $TARGET_DIR/$APP_NAME main.go

#go build -o $TARGET_DIR/$APP_NAME main.go
#go build -o $APP_NAME main.go
du -sh $TARGET_DIR/$APP_NAME
upx $TARGET_DIR/$APP_NAME
du -sh $TARGET_DIR/$APP_NAME
