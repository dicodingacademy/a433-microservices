#!/bin/bash

# build image dari dockerfile dengan nama item-app dan tag v1
echo "build docker"
docker build -t ghcr.io/footgraph/a433-microservices/karsajobs:latest .

# login ke githu container menggunakan personal access token yang sudah di set sebelum nya
echo $CR_PAT | docker login ghcr.io -u footgraph --password-stdin

# push image ke github container
docker push ghcr.io/footgraph/a433-microservices/karsajobs:latest