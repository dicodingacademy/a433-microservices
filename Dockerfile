# mengunduh/mengambil base image bernama node dari Docker Hub dengan tag 12-alpine
FROM node:14-bullseye

# membuat directory baru bernama app di dalam container dan menjadikannya sebagai working directory
WORKDIR /app

# menyalin semua berkas yang ada di local working directory saat ini ke container working directory (/app)
COPY . .

# menentukan agar aplikasi berjalan dalam production mode dan menggunakan container bernama item-dbsebagai database host
ENV NODE_ENV=production DB_HOST=item-db

# menginstal dependencies untuk production dan kemudian build aplikasi
RUN npm install --production --unsafe-perm && npm run build

# mengekspos port yang akan digunakan oleh container
EXPOSE 8080

# mengeksekusi perintah npm start untuk menjalankan aplikasi pada shell
CMD ["node", "start"]