#Membuat Docker image
docker build -t 12903478/karsajobs-ui:latest .

#Melihat daftar image
docker images

#Mengubah nama image sesuai format Docker Hub
docker tag item-app:v1 12903478/karsajobs-ui:latest

#Login ke Docker Hub
echo $PASS_DOCKER_HUB | docker login -u 12903478 --password-stdin

#Mengunggah image ke Docker Hub
docker push 12903478/karsajobs-ui:latest
