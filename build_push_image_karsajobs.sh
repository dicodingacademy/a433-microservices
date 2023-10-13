#!/bin/bash

# Gantilah nilai-nilai ini sesuai dengan kebutuhan Anda
GITHUB_USERNAME=d-oxys
GITHUB_REPOSITORY=a433-microservices  # Nama repositori GitHub Anda
IMAGE_NAME=karsajobs-backend
IMAGE_TAG=latest

# Nama lengkap Docker image yang akan dibangun
DOCKER_IMAGE=ghcr.io/$GITHUB_USERNAME/$GITHUB_REPOSITORY/$IMAGE_NAME:$IMAGE_TAG

# Build Docker image
docker build -t $DOCKER_IMAGE .

# Log in to GitHub Container Registry menggunakan token personal access
# Gantilah TOKEN_GITHUB_REGISTRY dengan token personal access Anda
docker login ghcr.io -u $GITHUB_USERNAME -p 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

# Push Docker image ke GitHub Container Registry
docker push $DOCKER_IMAGE
