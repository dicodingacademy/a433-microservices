# membuat Docker image
docker build -t item-app:v1 .

#melihat daftar image
docker images

#Mengubah nama image
docker tag item-app:v1 ghcr.io/bagus-k/item-app:v1

#save access token Github
export GITHUB_ACCESS_TOKEN="Your Token"

#Login ke GitHub Packages
echo $GITHUB_ACCESS_TOKEN | docker login ghcr.io -u bagusk --password-stdin 

#Mengunggah image Github Packages
docker push ghcr.io/bagus-k/item-app:v1