#!/bin/sh

# membuat variabel
export GH_REGISTRY=docker.pkg.github.com
export GH_REGISTRY_LOGIN=isutriadi
export GH_REGISTRY_REPO=$GH_REGISTRY/$GH_REGISTRY_LOGIN/containers
export GH_REGISTRY_IMAGE=karsajobs:latest

# melakukan build image
docker build -t $GH_REGISTRY_IMAGE .

# mengubah tag
docker tag localhost/$GH_REGISTRY_IMAGE $GH_REGISTRY_REPO/$GH_REGISTRY_IMAGE

# mengunggah image
docker push $GH_REGISTRY_REPO/$GH_REGISTRY_IMAGE
