#!/usr/bin/env bash

if [ -z $1 ]; then
	d="localhost:8080"
	go build ..
	# ../battlesnake-go > /dev/null &
	../battlesnake-go &
else
	d="$1"
fi

START=$(curl --silent -X POST -H \"Content-Type: application/json\" -d @../api/data/start.json $d/start)
echo "/start: $START"
MOVE=$(curl --silent -X POST -H \"Content-Type: application/json\" -d @../api/data/move.json $d/move)
echo "/move: $MOVE"

if [ -z $1 ]; then
	killall battlesnake-go > /dev/null
fi
