# Membuat Docker image dari Dockerfile
docker build -t ghcr.io/wahyurin/a433-microservices/karsajobs:latest .

# Melihat daftar image di lokal
docker images

# Login ke GitHub Container Registry
echo $TOKEN | docker login ghcr.io -u WahyuriN --password-stdin

# Mengunggah image ke GitHub Container Registry
docker push ghcr.io/wahyurin/a433-microservices/karsajobs:latest