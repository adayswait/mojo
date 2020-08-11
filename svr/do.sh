#!/bin/bash
if [ "$1" == "fmt" -o "$1" == "f" ]; then
	echo "formating all .go files"
	find . -path ./vendor -prune -o -print |grep '\.go' | xargs gofmt -w	
	echo "formating finished"
elif [ "$1" == "build" -o "$1" == "b" ]; then
	echo "go build -mod=vendor"
	go build -mod=vendor
	echo "building finished"
elif [ "$1" == "run" -o "$1" == "r" ]; then
	echo "go run -mod=vendor main.go"
	go run -mod=vendor main.go
	echo "building finished"
elif [ "$1" == "serve" -o "$1" == "s" ]; then
	echo "go run -mod=vendor main.go"
	go build -mod=vendor
	nohup ./mojo &
	echo "building finished"
else
	echo "usage:"
	echo ""
	echo "	do.sh fmt    formating all user .go codes"
	echo "	do.sh run    run this project"
	echo "	do.sh build  build this project"
	echo ""
	echo "	f short for fmt"
	echo "	r short for run"
	echo "	b short for build"
fi
