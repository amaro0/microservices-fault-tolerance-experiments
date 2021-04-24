#!/bin/bash
cd ..

docker build -f Dockerfile.finalserver -t finalserver .

for number in 1 2 3; do

  server="finalserver${number}"
  port="300${number}"
  echo "Initing ${server} on port: ${port}"

  docker stop $server
  docker rm $server
  docker run --name $server -d -p $port:3000 finalserver

done
