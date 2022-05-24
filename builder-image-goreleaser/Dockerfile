FROM alpine:latest AS goreleaser

LABEL org.opencontainers.image.source https://github.com/liatrio/builder-images

ARG GORELEASER_VERSION=v1.9.1
RUN wget https://github.com/goreleaser/goreleaser/releases/download/${GORELEASER_VERSION}/goreleaser_Linux_x86_64.tar.gz -O - | tar -xz

FROM golang:1.18-alpine

RUN apk add gcc musl-dev git

RUN addgroup -g 1000 jenkins && adduser -h /home/jenkins -G jenkins -u 1000 -D jenkins
USER jenkins

COPY --from=goreleaser goreleaser /usr/bin/goreleaser
