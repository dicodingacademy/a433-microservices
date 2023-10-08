#!/bin/bash

# Perintah untuk membuat Docker image dari Dockerfile
docker build -t ghcr.io/test-akun/karsajobs:latest .

# Login to GitHub Package
echo $CR_PAT | docker login ghcr.io -u test-akun --password-stdin


# Push image to GitHub Package
docker push ghcr.io/test-akun/karsajobs:latest
