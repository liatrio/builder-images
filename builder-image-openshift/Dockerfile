FROM alpine:3.9

LABEL org.opencontainers.image.source https://github.com/liatrio/builder-images

RUN apk update && \
    apk add curl && \
    apk add --no-cache libc6-compat && \
    apk add ca-certificates

RUN wget https://github.com/openshift/origin/releases/download/v3.11.0/openshift-origin-client-tools-v3.11.0-0cbc58b-linux-64bit.tar.gz -O - | tar -xz
ENV PATH ${PATH}:/openshift-origin-client-tools-v3.11.0-0cbc58b-linux-64bit


ENV HELM_VERSION 2.13.1
RUN curl -f -L https://get.helm.sh/helm-v${HELM_VERSION}-linux-amd64.tar.gz | tar -zx && \
  chmod +x linux-amd64/helm && \
  mv linux-amd64/helm /usr/bin && \
  rm -rf linux-amd64
