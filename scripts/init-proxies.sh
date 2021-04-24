#!/bin/bash
cd ..

docker build -f Dockerfile.proxyserver -t proxyserver .

for number in 0 1 2; do
  server="proxyserver${number}"
  port="400${number}"
  echo "Initing ${server} on port: ${port}"

  docker stop $server
  docker rm $server
  docker run --name $server -d -p $port:4000 proxyserver

done
