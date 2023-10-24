#!/bin/bash

# Build Docker image
docker build -t ghcr.io/irvanfebri/shipping-service:latest .

# Log in to GitHub Container Registry
echo $GH_PACKAGES_TOKEN | docker login ghcr.io -u irvanfebri --password-stdin

# Push Docker image to GitHub Container Registry
docker push ghcr.io/irvanfebri/shipping-service:latest
