# syntax=docker/dockerfile:1.7-labs

FROM node:22-alpine
WORKDIR /app

COPY --exclude=node_modules . .

RUN npm install

EXPOSE 5173

CMD ["npm", "run", "dev" ]
