FROM node:6.9.4-alpine
RUN mkdir -p /usr/src/app
COPY package.json /usr/src/app/
RUN npm install
COPY . /usr/src/app
CMD ["npm","start"]