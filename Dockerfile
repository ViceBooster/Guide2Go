FROM debian:10-slim

ENV PORT=3000
ENV IP_ADDRESS="192.168.200.1"

COPY guide2go /usr/local/bin/guide2go

RUN apt-get update && apt-get install ca-certificates -y && apt autoclean

ENTRYPOINT ["tail", "-f", "/dev/null"]


