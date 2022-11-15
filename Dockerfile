# syntax=docker/dockerfile:1

FROM golang:1.19-bullseye

WORKDIR /image
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.50.1
ADD requirements.txt .
RUN DEBIAN_FRONTEND="noninteractive" apt-get update && apt-get -y install python3-pip docker.io && pip install -r requirements.txt
RUN apt-get autoremove -y \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/* \
    && rm -rf /var/cache/* \
    && rm -rf /var/log/apt\
    && rm -rf /root/.cache
