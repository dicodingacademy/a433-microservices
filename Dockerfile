# Node.js version 14 base image
FROM node:14
# Production app runs on port 8080
EXPOSE 8080
# Set production mode and use app-db as the database host
ENV NODE_ENV=production DB_HOST=app-db 
# Set working directory to where source is
WORKDIR /app
# Copy source files into container
COPY . /app
# Install production dependencies and build app
RUN npm install --production --unsafe-perm && npm run build
# Start the server in production mode and use app-db as db host
CMD ["npm", "start"]