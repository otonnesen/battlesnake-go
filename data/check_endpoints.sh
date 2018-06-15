#!/usr/bin/env bash
START=$(curl --silent -X POST -H \"Content-Type: application/json\" -d @data/start_request.json http://localhost:8080/start)
echo "/start: $START"
MOVE=$(curl --silent -X POST -H \"Content-Type: application/json\" -d @data/move_request.json http://localhost:8080/move)
echo "/move: $MOVE"
END=$(curl --silent -X POST -H \"Content-Type: application/json\" -d @data/end_request.json http://localhost:8080/end)
echo "/end: $END"
