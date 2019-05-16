#FROM alpine:3.9
FROM docker:18.09

RUN apk add --no-cache \
		git \
		curl \
		openssh-client \
		make

ENV SKAFFOLD_VERSION 0.29.0
RUN curl -f -Lo skaffold https://github.com/GoogleCloudPlatform/skaffold/releases/download/v${SKAFFOLD_VERSION}/skaffold-linux-amd64 && \
  chmod +x skaffold && \
  mv skaffold /usr/bin

ENV CST_VERSION 1.8.0
RUN curl -f -Lo container-structure-test https://storage.googleapis.com/container-structure-test/v${CST_VERSION}/container-structure-test-linux-amd64 && \
  chmod +x container-structure-test && \
  mv container-structure-test /usr/bin

ENV HELM_VERSION 2.13.1
RUN curl -f -L https://storage.googleapis.com/kubernetes-helm/helm-v${HELM_VERSION}-linux-amd64.tar.gz | tar -zx && \
  chmod +x linux-amd64/helm && \
  mv linux-amd64/helm /usr/bin && \
  rm -rf linux-amd64
