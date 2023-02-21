FROM node:19-alpine as build

WORKDIR /app

COPY . .

RUN yarn install

EXPOSE 4200

CMD ["yarn", "dev"]
