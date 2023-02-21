#membuat Docker image dari Dockerfile
docker build -t item-app:v1 .

#melihat daftar image di lokal
docker image ls

#mengubah nama image sesuai dengan format Github Packages
docker tag item-app:v1 ghcr.io/trisnanto/item-app:v1

#login ke Github Packages
echo $CR_PAT | docker login ghcr.io -u trisnanto --password-stdin

#mengunggah image ke Github Packages
docker push ghcr.io/trisnanto/item-app:v1
