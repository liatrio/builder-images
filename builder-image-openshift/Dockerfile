FROM alpine:3.15

LABEL org.opencontainers.image.source https://github.com/liatrio/builder-images

RUN apk update && \
    apk add curl && \
    apk add --no-cache libc6-compat && \
    apk add ca-certificates

RUN wget https://github.com/openshift/origin/releases/download/v3.11.0/openshift-origin-client-tools-v3.11.0-0cbc58b-linux-64bit.tar.gz -O - | tar -xz
ENV PATH ${PATH}:/openshift-origin-client-tools-v3.11.0-0cbc58b-linux-64bit

COPY --from=alpine/helm:3.9.0 /usr/bin/helm /usr/bin/
