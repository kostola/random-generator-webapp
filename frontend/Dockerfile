# Stage 1: Build the Svelte app
FROM node:lts as build
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY ./src ./src
COPY ./static ./static
COPY .env .
COPY svelte.config.js .
COPY vite.config.js .
RUN ls -la
RUN npm run build
RUN ls -la build

# Stage 2: Create a plain Node container to run the app
FROM node:lts-alpine
WORKDIR /app
COPY --from=build /app/build ./build
COPY --from=build /app/package.json .
EXPOSE 3000
ENV PUBLIC_API_BASE_URL="http://localhost:8081"
CMD [ "node", "build/index.js" ]
