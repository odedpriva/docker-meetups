FROM node:8-alpine
RUN mkdir -p /usr/src/app
WORKDIR  /usr/src/app/
COPY package.json /usr/src/app/
RUN npm install
COPY . /usr/src/app
ARG COMMIT
ENV COMMIT ${COMMIT:-no-commit}
LABEL version="1.0"
ONBUILD RUN echo this is onbuild command >> /tmp/onbuild
CMD ["npm","start"]