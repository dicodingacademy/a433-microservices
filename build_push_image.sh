#!/bin/bash

# melakukan build image dengan tag menggunakan file Dockerfile
docker build -t item-app:v1 .

# mengecek docker images yang berada di lokal
docker images

# melakukan perubahan tag agar sesuai dengan format github container registry / ghcr.io
docker tag item-app:v1 ghcr.io/man20820/item-app:v1

# melakukan login ke github container registry
# echo $PASSWORD_DOCKER_HUB | docker login ghcr.io -u man20820 --password-stdin

# melakukan push image ke github container registry
docker push ghcr.io/man20820/item-app:v1