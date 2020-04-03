FROM golang:1.14.1 as builder

WORKDIR /usr/src

RUN git clone https://github.com/mrparkers/skaffold.git && \
    cd skaffold && \
    GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
        go build -tags "osusergo netgo static_build release" \
        -ldflags " -X github.com/GoogleContainerTools/skaffold/pkg/skaffold/version.version=v1.6.0-55-g145f59579 -X github.com/GoogleContainerTools/skaffold/pkg/skaffold/version.gitCommit=145f59579470eb1f0a7f40d8e0924f8716c6f05b -X github.com/GoogleContainerTools/skaffold/pkg/skaffold/version.gitTreeState=clean -s -w  -extldflags \"-static\"" \
        -o out/skaffold github.com/GoogleContainerTools/skaffold/cmd/skaffold

FROM docker:18.09

RUN apk add --no-cache \
		git \
		curl \
		openssh-client \
		make

#ENV SKAFFOLD_VERSION 1.6.0
#RUN curl -f -Lo skaffold https://github.com/GoogleCloudPlatform/skaffold/releases/download/v${SKAFFOLD_VERSION}/skaffold-linux-amd64 && \
#  chmod +x skaffold && \
#  mv skaffold /usr/bin

COPY --from=builder /usr/src/skaffold/out/skaffold /usr/bin/skaffold
RUN skaffold version

ENV CST_VERSION 1.8.0
RUN curl -f -Lo container-structure-test https://storage.googleapis.com/container-structure-test/v${CST_VERSION}/container-structure-test-linux-amd64 && \
  chmod +x container-structure-test && \
  mv container-structure-test /usr/bin

ENV HELM_VERSION 3.1.2
RUN curl -f -L https://get.helm.sh/helm-v${HELM_VERSION}-linux-amd64.tar.gz | tar -zx && \
  chmod +x linux-amd64/helm && \
  mv linux-amd64/helm /usr/bin && \
  rm -rf linux-amd64
