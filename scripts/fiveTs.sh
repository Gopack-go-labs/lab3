#!/bin/bash

initCoord=5
finalLeftCoord=2
finalRighttCoord=8
timeinterval=0.1
step=1

sleep $timeinterval

curl -X POST http://localhost:17000 -d "reset"
curl -X POST http://localhost:17000 -d "update"

while true; do

curl -X POST http://localhost:17000 -d "white"
curl -X POST http://localhost:17000 -d "green"

sleep $timeinterval

curl -X POST http://localhost:17000 -d "figure 0.$initCoord 0.$initCoord"
sleep $timeinterval
curl -X POST http://localhost:17000 -d "figure 0.$finalLeftCoord 0.$finalLeftCoord"
sleep $timeinterval
curl -X POST http://localhost:17000 -d "figure 0.$finalRighttCoord 0.$finalRighttCoord"
sleep $timeinterval
curl -X POST http://localhost:17000 -d "figure 0.$finalRighttCoord 0.$finalLeftCoord"
sleep $timeinterval
curl -X POST http://localhost:17000 -d "figure 0.$finalLeftCoord 0.$finalRighttCoord"
sleep $timeinterval

curl -X POST http://localhost:17000 -d "update"

curl -X POST http://localhost:17000 -d "reset"
curl -X POST http://localhost:17000 -d "update"
done

