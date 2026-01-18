FROM ubuntu:24.04

### go environment
RUN apt update && apt install -y \
    ca-certificates \
    lsb-release \
    curl \
    git \
    iputils-ping

COPY --from=golang:1.25.5-alpine3.23 /usr/local/go /usr/local/go
ENV PATH="$PATH:/usr/local/go/bin"
WORKDIR /opt/Dev