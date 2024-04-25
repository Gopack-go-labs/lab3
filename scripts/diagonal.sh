#!/bin/bash

initCoord=5
finalLeftCoord=2
finalRighttCoord=8
interval=0.1
step=1

x=$initCoord
y=$initCoord

curl -X POST http://localhost:17000 -d "white"

curl -X POST http://localhost:17000 -d "figure 0.$initCoord 0.$initCoord"

sleep $interval

  while ((x > finalLeftCoord)); do

    x=$((x - step))
    curl -X POST http://localhost:17000 -d "move 0.$x 0.$x "
    curl -X POST http://localhost:17000 -d "update"
    sleep $interval
    
  done

  while ((x < finalRighttCoord)); do

    x=$((x + step))
    curl -X POST http://localhost:17000 -d "move 0.$x 0.$x "
    curl -X POST http://localhost:17000 -d "update"
    sleep $interval
    
  done


  while ((x > initCoord)); do

    x=$((x - step))
    curl -X POST http://localhost:17000 -d "move 0.$x 0.$x "
    curl -X POST http://localhost:17000 -d "update"
    sleep $interval
    
  done


  while ((x < finalRighttCoord)); do

    x=$((x + step))
    y=$((y - step))
    curl -X POST http://localhost:17000 -d "move 0.$x 0.$y "
    curl -X POST http://localhost:17000 -d "update"
    sleep $interval
  done

  while ((x > finalLeftCoord)); do

    x=$((x - step))
    y=$((y + step))
    curl -X POST http://localhost:17000 -d "move 0.$x 0.$y "
    curl -X POST http://localhost:17000 -d "update"
    sleep $interval
  done

  while ((x < initCoord)); do

    x=$((x + step))
    y=$((y - step))
    curl -X POST http://localhost:17000 -d "move 0.$x 0.$y "
    curl -X POST http://localhost:17000 -d "update"
    sleep $interval
  done


sleep $interval

curl -X POST http://localhost:17000 -d "green"

curl -X POST http://localhost:17000 -d "update"