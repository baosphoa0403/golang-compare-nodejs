FROM node:23.8.0

WORKDIR /usr/src/app

COPY package*.json ./


RUN npm install

COPY . .

EXPOSE 3000

CMD ["node", "app.js"]
