FROM alpine:3.6
RUN apk add --no-cache \
    curl \
    tcpdump \
    iptables \
    drill \
    iproute2 \
    ebtables \
    ethtool \
    redis \
    postgresql-client \
    tshark
RUN apk add --no-cache \ 
    bash \
    bash-completion
RUN curl -L https://github.com/sequenceiq/docker-alpine-dig/releases/download/v9.10.2/dig.tgz|tar -xzv -C /usr/local/bin/
COPY nsenter /bin/nsenter
RUN chmod +x /bin/nsenter
RUN apk add --update netcat-openbsd && rm -rf /var/cache/apk/*