#!/bin/bash

# Конфігурація
server_port=17000
initCoord=5
finalLeftCoord=2
finalRighttCoord=8
timeinterval=0.1
step=1

x=$initCoord
y=$initCoord

curl -X POST http://localhost:17000 -d "reset"
curl -X POST http://localhost:17000 -d "update"

function move_figure() {
  curl -X POST http://localhost:$server_port -d "move 0.$x 0.$y"
  curl -X POST http://localhost:$server_port -d "update"
  sleep $timeinterval
}

curl -X POST http://localhost:17000 -d "white"

curl -X POST http://localhost:17000 -d "figure 0.$initCoord 0.$initCoord"

sleep $timeinterval

while ((x > finalLeftCoord)); do
  x=$((x - step))
  y=$((y - step))
  move_figure
done

while ((x < finalRighttCoord)); do
  x=$((x + step))
  y=$((y + step))
  move_figure
done

while ((x > initCoord)); do
  x=$((x - step))
  y=$((y - step))
  move_figure
done

while ((x < finalRighttCoord)); do
  x=$((x + step))
  y=$((y - step))
  move_figure
done

while ((x > finalLeftCoord)); do
  x=$((x - step))
  y=$((y + step))
  move_figure
done

while ((x < initCoord)); do
  x=$((x + step))
  y=$((y - step))
  move_figure
done

sleep $timeinterval

curl -X POST http://localhost:17000 -d "green"

curl -X POST http://localhost:17000 -d "update"