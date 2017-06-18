FROM node:6.9.4-alpine
RUN mkdir -p /usr/src/app
COPY . /usr/src/app
COPY package.json /usr/src/app/
RUN npm install
CMD ["npm","start"]