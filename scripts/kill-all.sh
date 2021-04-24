#!/bin/bash
for number in 0 1 2; do
  finalserver="finalserver${number}"
  proxyserver="proxyserver${number}"

  docker stop $finalserver
  docker rm $finalserver
  docker stop $proxyserver
  docker rm $proxyserver
done
