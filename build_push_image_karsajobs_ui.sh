#!/bin/sh

# build Docker image dari berkas Dockerfile yang disediakan
docker build -t karsajobs-ui .

# Mengubah nama image agar sesuai dengan format Github Packages
docker tag karsajobs-ui ghcr.io/rickybeatus/karsajobs-ui

# Login ke Github Packages
echo $CR_PAT | docker login ghcr.io -u rickybeatus --password-stdin

# Mengunggah image ke Github Packages
docker push ghcr.io/rickybeatus/karsajobs-ui