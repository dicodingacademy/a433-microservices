# Menggunakan image Node.js versi 18 yang berbasis Alpine Linux sebagai builder
FROM node:18-alpine as builder

# Menentukan direktori kerja di dalam container
WORKDIR /app

# Menyalin package.json dan package-lock.json ke direktori kerja
COPY package*.json ./

# Menjalankan npm ci untuk menginstal dependensi dengan package-lock.json
RUN npm ci

# Menyalin semua berkas JavaScript dari direktori saat ini ke direktori kerja
COPY ./*.js ./

# Mengekspos port 3000 yang akan digunakan
EXPOSE 3000

# Menjalankan perintah "node index.js" saat container berjalan
CMD [ "node", "index.js" ]
