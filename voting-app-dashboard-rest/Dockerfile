# FROM node:8.9
# FROM grpc/node
# FROM node:8.9-alpine
# FROM node:alpine
# FROM mhart/alpine-node:10
FROM node:10-alpine

# RUN apk add --no-cache make gcc g++ python/
RUN mkdir -p /app
WORKDIR /app

# RUN apk update && apk upgrade
# RUN apk add --no-cache libstdc++ python make g++

RUN npm install -g nodemon
RUN npm config set registry https://registry.npmjs.org
# COPY package.json /app/package.json
COPY . /app
RUN ls -al .
RUN rm -rf node_modules && rm package-lock.json
RUN ls -al .
RUN npm install
# RUN npm cache clean --force
# RUN mv /app/node_modules /node_modules
# RUN npm ls grpc
# RUN ls -al node_modules/grpc/src/node/extension_binary/
# COPY . /app
# RUN npm install grpc
# RUN npm rebuild

# RUN npm rebuild --target=8.1.0 --target_platform=linux --target_arch=x64 --target_libc=musl --update-binary
# RUN npm rebuild --target=9.0.0 --target_libc=musl --target_platform=linux
# RUN npm rebuild --target=8.1.0 --target_platform=linux --target_arch=x64 --target_libc=glibc --update-binary
# RUN npm rebuild --target=8.1.0 --target_platform=linux --target_arch=x64 --target_libc=musl --update-binary
RUN ls -al .
RUN ls -al node_modules/grpc/src/node/extension_binary/

EXPOSE 8081

CMD ["nodemon", "server.js"]
