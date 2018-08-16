#!/usr/bin/env bash
START=$(curl --silent -X POST -H \"Content-Type: application/json\" -d @start_request.json $1/start)
echo "/start: $START"
MOVE=$(curl --silent -X POST -H \"Content-Type: application/json\" -d @move_request.json $1/move)
echo "/move: $MOVE"
END=$(curl --silent -X POST -H \"Content-Type: application/json\" -d @end_request.json $1/end)
echo "/end: $END"
