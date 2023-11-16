# Membuat Docker image dari Dockerfile sesuai tag Github Container Registry
docker build -t ghcr.io/yahyakhaliman/order-service . 

# Login ke Github Container Registry
echo $PAT | docker login ghcr.io --username yahyakhaliman --password-stdin 

# Mengunggah image ke Github Container Registry
docker push ghcr.io/yahyakhaliman/order-service