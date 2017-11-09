FROM node:alpine
# everything in container related to app directory
WORKDIR /app
COPY package.json .
RUN npm install
COPY . .
EXPOSE 80
ENTRYPOINT ["node", "index.js"]