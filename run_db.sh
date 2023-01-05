#!/bin/bash

echo "Opening Docker"
open -a docker
sleep 10
echo "remove Old Docker image for postgres"
docker rm -postgres
sleep 10
echo "Create new postgres"

docker run --name postgres -p 127.0.0.1:5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=postgres -d postgres
