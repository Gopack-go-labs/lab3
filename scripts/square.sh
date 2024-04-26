#!/bin/bash

# Конфігурація
server_port=17000
initCoord=5
finalLeftCoord=2
finalRightCoord=8
timeinterval=0.1
step=1

# Очищення
curl -X POST http://localhost:17000 -d "reset"
curl -X POST http://localhost:17000 -d "update"

# Початкове розташування фігури
curl -X POST http://localhost:$server_port -d "figure 0.$finalLeftCoord 0.$finalLeftCoord"

# Функція для переміщення фігури
function move_figure() {
  local x=$1
  local y=$2
  curl -X POST http://localhost:$server_port -d "move 0.$x 0.$y"
  curl -X POST http://localhost:$server_port -d "update"
  sleep $timeinterval
}

# Переміщення фігури в квадраті
while true; do
  x=$finalLeftCoord
  y=$finalLeftCoord

  # Переміщення вправо
  while ((x < finalRightCoord)); do
    x=$((x + step))
    move_figure $x $y
  done

  # Переміщення вниз
  while ((y < finalRightCoord)); do
    y=$((y + step))
    move_figure $x $y
  done

  # Переміщення вліво
  while ((x > finalLeftCoord)); do
    x=$((x - step))
    move_figure $x $y
  done

  # Переміщення вверх
  while ((y > finalLeftCoord)); do
    y=$((y - step))
    move_figure $x $y
  done

  # Оновлення фігури
  curl -X POST http://localhost:$server_port -d "update"
done