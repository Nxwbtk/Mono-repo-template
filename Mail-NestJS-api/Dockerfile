FROM node:23-bookworm-slim

WORKDIR /usr/src/app

COPY --chown=node:node app/package*.json ./

RUN npm ci

COPY --chown=node:node app/. ./

RUN npm run build

EXPOSE 7652

CMD [ "node", "dist/main.js" ]