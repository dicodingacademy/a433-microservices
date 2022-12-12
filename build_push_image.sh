#!/bin/sh

# membuat sebuah Docker image menggunakan Dockerfile yang tadi dibuat dengan nama item-app dan tag v1
docker build -t item-app:v1 .

# melihat daftar image di lokal
docker images

# mengubah nama image agar sesuai dengan format Github Packages
docker tag item-app:v1 filahaa/item-app:v1

# login ke Github Packages untuk keamanan saya menggunakan Token
echo $TOKEN_GITHUB | docker login ghcr.io -u filahaa --password-stdin

# mengunggah image ke Github Packages
docker push ghcr.io/filahaa/item-app:v1