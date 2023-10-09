FROM node:18.18.0-alpine3.18 AS build

WORKDIR /app

COPY package.json package-lock.json ./
RUN npm ci
COPY . .
RUN npm run build

FROM node:18.18.0-alpine3.18 AS prod

USER node:node
WORKDIR /app

COPY --from=build --chown=node:node /app/build ./build
COPY --from=build --chown=node:node /app/package.json ./

CMD ["node", "build"]