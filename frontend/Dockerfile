FROM node:alpine AS build-stage

WORKDIR /usr/local/app

COPY . /usr/local/app

RUN npm install && \
    npm run build && \
    npm install -g serve && \
    rm -rf src && \
    npm prune -omit=dev && \
    npm cache clean --force && \
    rm -rf /var/lib/apt/lists/*

FROM node:alpine 

WORKDIR /usr/local/app

EXPOSE 5001

COPY --from=build-stage /usr/local/app/dist ./dist

RUN npm install -g serve && adduser -D appuser && chown appuser /usr/local/app

USER appuser

ENTRYPOINT [ "serve", "-s", "-l", "5001", "dist"]