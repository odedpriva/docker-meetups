FROM ubuntu:16.04
RUN apt update && \ 
rm -rf /var/lib/apt/lists/*