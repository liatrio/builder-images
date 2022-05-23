FROM alpine:3.15

LABEL org.opencontainers.image.source https://github.com/liatrio/builder-images

RUN apk add --no-cache \
    groff \
    python3 \
    py3-pip

RUN pip3 install awscli
