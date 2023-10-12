# Stage 1: Build the Svelte app
FROM node:lts as build
# Set the working directory
WORKDIR /app
# Copy the package.json and package-lock.json
COPY package*.json ./
# Install dependencies
RUN npm install
# Copy the Svelte app source code
COPY . .
RUN ls -la
# Build the Svelte app
RUN npm run build
RUN ls -la build

# Stage 2: Create a plain Node container to run the app
FROM node:lts
# Set the working directory in the final image
WORKDIR /app
# Copy the built Svelte app from the previous stage
COPY --from=build /app/build .
# Expose the port the app will run on (adjust as needed)
EXPOSE 3000
ENV PUBLIC_API_BASE_URL="http://localhost:8081"
# Command to run the app
CMD [ "node", "index.js" ]
