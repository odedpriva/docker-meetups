FROM ubuntu:17.04
RUN apt-get update && apt-get install -y \
    curl \
 && rm -rf /var/lib/apt/lists/*
RUN curl -O https://storage.googleapis.com/golang/go1.8.linux-amd64.tar.gz
RUN tar -xvf go1.8.linux-amd64.tar.gz
RUN mv go /usr/local
RUN echo 'export PATH=$PATH:/usr/local/go/bin' >> /etc/bash.bashrc
RUN echo 'export GOROOT=/usr/local/go' >> /etc/bash.bashrc
RUN echo 'export PATH=$PATH:$GOROOT/bin' >> /etc/bash.bashrc

