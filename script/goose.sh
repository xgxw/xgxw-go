#!/bin/bash

# go get -u github.com/pressly/goose/cmd/goose

goose mysql "root:test@/xgxw?parseTime=true" up

goose mysql "root:test@/xgxw?parseTime=true" status

