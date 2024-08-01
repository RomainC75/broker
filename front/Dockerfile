# FROM node:20.14-alpine3.20
FROM node:18.20.4-alpine
WORKDIR /app

COPY . ./
RUN npm install && npm cache clean --force

COPY tsconfig.json ./

EXPOSE 5173

CMD ["npm", "run", "dev"]