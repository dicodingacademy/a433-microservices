# Menggunakan base image Node.js versi 14
FROM node:14

# Menentukan working directory untuk container
WORKDIR /app

# Menyalin seluruh source code ke working directory di container
COPY . .

# Menentukan variabel lingkungan untuk Node.js dan database host
ENV NODE_ENV=production DB_HOST=item-db

# Menginstal dependencies untuk production dan kemudian build aplikasi
RUN npm install --production --unsafe-perm && npm run build

# Ekspos port 8080 yang digunakan oleh aplikasi
EXPOSE 8080

# Saat container diluncurkan, jalankan server dengan perintah npm start
CMD ["npm", "start"]
