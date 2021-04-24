#!/bin/bash

sh init-proxies.sh
sh init-finals.sh

go run ../metrics.go
