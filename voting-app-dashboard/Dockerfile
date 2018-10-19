# FROM node:8.9
FROM grpc/node:1.0-onbuild

# RUN apk add --no-cache make gcc g++ python/
RUN mkdir -p /app
WORKDIR /app

RUN npm install -g nodemon
RUN npm config set registry https://registry.npmjs.org
COPY package.json /app/package.json
RUN npm install
RUN npm cache clean --force
RUN mv /app/node_modules /node_modules
COPY . /app

ENV PORT 8081
EXPOSE 8081

CMD ["node", "server.js"]
