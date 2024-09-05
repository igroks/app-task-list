FROM node:19-alpine

WORKDIR /app

COPY . .

EXPOSE 4200

RUN yarn install

CMD [ "yarn", "dev"]
