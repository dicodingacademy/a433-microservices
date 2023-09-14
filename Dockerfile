# menggunakan nodejs versi 14
FROM node:14

# mengubah workdir ke /app
WORKDIR /app

# melakukan copy directory saat ini (mesin) ke directory workdir (/app di dalam image)
COPY . .

# menyetting environment variables di image
ENV NODE_ENV=production DB_HOST=item-db

# install kebutuhan aplikasi & build
RUN npm install --production --unsafe-perm && npm run build

# menginformasikan port 8080 yang digunakan
EXPOSE 8080

# menjalankan secara default ketika container berjalan
CMD [ "npm", "start" ]