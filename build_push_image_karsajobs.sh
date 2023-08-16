#!/bin/sh

# build image docker dari Dockerfile dari current working directory dengan nama hann315/karsajobs dan tag lastest
docker build -t hann315/karsajobs:latest .
# membuat line baru
echo '\n\n'

# membuat tag image baru dari source image menggunakan image hann315/karsajobs:latest, tetapi dengan nama/tag yang beda
docker tag hann315/karsajobs:latest ghcr.io/hann315/karsajobs:latest
# membuat line baru
echo '\n\n'

# Melakukan export token github packages ke environment variable
export GITHUB_PACKAGES_TOKEN=

# login ke github packages
echo $GITHUB_PACKAGES_TOKEN | docker login ghcr.io -u hann315 --password-stdin
# membuat line baru
echo '\n\n'

# push image ke github packages
docker push ghcr.io/hann315/karsajobs:latest
# membuat line baru
echo '\n\n'