FROM node:lts-alpine3.19
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
# RUN npm run build
EXPOSE 3000
CMD ["node", "index.js"]