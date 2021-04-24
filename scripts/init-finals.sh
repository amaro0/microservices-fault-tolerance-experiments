#!/bin/bash
cd ..

docker build -f Dockerfile.finalserver -t finalserver .

for number in 0 1 2; do
  server="finalserver${number}"
  port="300${number}"
  echo "Initing ${server} on port: ${port}"

  docker stop $server
  docker rm $server

  if [ -f .env ]
  then
    docker run --name $server -d -p $port:3000 --env-file .env finalserver
  else
    docker run --name $server -d -p $port:3000 finalserver
  fi

done
