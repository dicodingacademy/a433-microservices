# menggunakan node versi 14 alpine
FROM node:14.21.2-alpine

# Menentukan working directory container adalah /app
WORKDIR /app

# Menyalin seluruh source code ke working directory di container
COPY . .

#Menentukan agar aplikasi berjalan dalam production mode
ENV NODE_ENV=production DB_HOST=item-db

#Menginstal dependencies untuk production & Build
RUN npm install --production --unsafe-perm && npm run build

#Expose port ke 8080
EXPOSE 8080

#run npm start
CMD ["npm","start"]
