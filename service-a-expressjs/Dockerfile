FROM node:20.15.0 AS build

WORKDIR /app

COPY package.json yarn.lock ./

RUN yarn install

COPY . ./

FROM node:20.15.0-alpine3.20

WORKDIR /

COPY --from=build /app ./

EXPOSE 3000

ENTRYPOINT ["yarn", "start"]
