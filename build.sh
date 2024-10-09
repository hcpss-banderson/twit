#!/usr/bin/env bash

archs=(amd64 arm64)

for arch in ${archs[@]}
do
	env GOOS=linux GOARCH=${arch} go build -o twit-${arch}
done
