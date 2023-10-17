#!/bin/bash

# Build Docker image
docker build -t ghcr.io/footgraph/shipping-service:latest .

# export dockerhub access_token
export ACCESS_TOKEN_DOCKER_HUB=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

# Log in to GitHub Container Registry
echo $ACCESS_TOKEN_DOCKER_HUB | docker login ghcr.io -u footgraph --password-stdin

# Push Docker image to GitHub Container Registry
docker push ghcr.io/footgraph/shipping-service:latest