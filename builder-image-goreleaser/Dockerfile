FROM alpine:latest AS goreleaser

RUN wget https://github.com/goreleaser/goreleaser/releases/download/v0.123.3/goreleaser_Linux_x86_64.tar.gz -O - | tar -xz

FROM golang:1.13-alpine

RUN apk add gcc musl-dev git

RUN addgroup -g 1000 jenkins && adduser -h /home/jenkins -G jenkins -u 1000 -D jenkins
USER jenkins

COPY --from=goreleaser goreleaser /usr/bin/goreleaser
