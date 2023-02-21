# menggunakan base image Node.js ver 14
FROM node:14-alpine

# menambahkan metadata kedalam image
LABEL org.opencontainers.image.source=https://github.com/trisnanto/a433-microservices

# menentukan working directory didalam container
WORKDIR /app

# menyalin seluruh source code ke working directory didalam container
COPY . .

# mengatur variabel environment agar aplikasi berjalan dalam production mode
# dan menggunakan container bernama item-db sebagai database host
ENV NODE_ENV=production DB_HOST=item-db

# menginstall dependencies untuk production dan kemudian build aplikasi
RUN npm install --production --unsafe-perm && npm run build

# menggunakan port 8080 untuk aplikasi
EXPOSE 8080

# menjalankan server
CMD ["npm", "start"]