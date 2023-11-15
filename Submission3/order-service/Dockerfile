# Menggunakan base image Node.js versi 14
FROM node:14 

# Menentukan working directory untuk container
WORKDIR /app

# Salin package.json dan package-lock.json ke direktori kerja
COPY package*.json ./

# Menginstal dependencies yang diperlukan
RUN npm install

# Menyalin seluruh source code ke working directory di container
COPY . .

# Ekspos port 8080 yang digunakan oleh aplikasi
EXPOSE 3000

# Set Environment Variabel
ENV PORT=3000
ENV AMQP_URL="amqp://localhost:5672"

# Saat container diluncurkan, jalankan server dengan perintah npm start
CMD ["npm", "run", "dev"]
