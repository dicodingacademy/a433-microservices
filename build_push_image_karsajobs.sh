#!/bin/bash

# Build image dengan tag "ghcr.io/dionanda/karsajobs:latest" menggunakan Dockerfile pada direktori saat ini
docker build -t ghcr.io/dionanda/karsajobs:latest .

# Melihat daftar image Docker yang ada
docker images

# Menggunakan environment variable $PAT untuk login ke GitHub Container Registry
read -r PAT
echo $PAT | docker login ghcr.io -u dionanda --password-stdin

# Mengunggah image dengan tag "ghcr.io/dionanda/karsajobs:latest" ke GitHub Container Registry
docker push ghcr.io/dionanda/karsajobs:latest