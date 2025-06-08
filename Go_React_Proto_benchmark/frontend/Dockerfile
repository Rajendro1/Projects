FROM node:24 AS proto-gen

WORKDIR /app

RUN apt-get update && apt-get install -y git python3 make g++ && rm -rf /var/lib/apt/lists/*

# Install protobufjs CLI package
RUN npm install protobufjs-cli

COPY proto/user.proto ./proto/

RUN mkdir -p src/proto && \
    npx pbjs -t static-module -w commonjs -o src/proto/user_pb.js proto/user.proto && \
    npx pbts -o src/proto/user_pb.d.ts src/proto/user_pb.js


# Stage 2: Build React app with Vite
FROM node:24-alpine AS build

WORKDIR /app

COPY frontend/package*.json ./
RUN npm install

# Copy generated proto files from proto-gen stage
COPY --from=proto-gen /app/src/proto ./src/proto

# Copy React source code
COPY frontend/ ./

RUN npm run build

RUN ls -la /app/dist


# Stage 3: Serve with nginx
FROM nginx:alpine

# Copy React build output (dist) to nginx html folder
COPY --from=build /app/dist /usr/share/nginx/html

# Expose port 80
EXPOSE 80

# Start nginx
CMD ["nginx", "-g", "daemon off;"]
