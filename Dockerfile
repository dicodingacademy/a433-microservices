# import node v14 on alpine
FROM node:14-alpine

# deklarasi workdir
WORKDIR /app

# copy semua file ke workdir
COPY . .

# install dependencies yang dibutuhkan
RUN npm install

# deklarasi environtment yang dibutuhkan
ENV PORT=3000

# expose port 3001
EXPOSE 3000

# run npm start jika kontainer diluncurkan
CMD ["npm","start"]