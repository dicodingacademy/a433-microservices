# Use Node.js version 14 as the base image
FROM node:14

# Set the working directory inside the container to /src
WORKDIR /src

# Copy package.json and package-lock.json to the working directory
COPY package*.json ./

# Install dependencies using npm
RUN npm install

# Download the wait-for-it.sh script from the repository
RUN wget -O ./wait-for-it.sh https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh

# Make the wait-for-it.sh script executable
RUN chmod +x ./wait-for-it.sh

# Copy the index.js file to the working directory
COPY index.js ./

# Set the environment variable PORT to 3001
ENV PORT=3001

# Expose the specified port
EXPOSE $PORT

# Define the command to run the application after waiting for RabbitMQ to start
CMD ["sh", "-c", "./wait-for-it.sh my-rabbitmq:5672 --timeout=30 -- node index.js"]