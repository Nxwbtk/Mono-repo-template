FROM node:24.5-alpine

WORKDIR /usr/src/app

COPY --chown=node:node app/package*.json ./

RUN npm ci

COPY --chown=node:node app/. ./

RUN npm run build

EXPOSE 7652

CMD [ "node", "dist/main.js" ]