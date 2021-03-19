FROM docker:18.09-dind

RUN apk add --no-cache \
		git \
		curl \
		openssh-client \
		make \
		groff \
		python3

ENV SKAFFOLD_VERSION 1.15.0
RUN curl -f -Lo skaffold https://github.com/GoogleCloudPlatform/skaffold/releases/download/v${SKAFFOLD_VERSION}/skaffold-linux-amd64 && \
  chmod +x skaffold && \
  mv skaffold /usr/bin && \
  skaffold version

RUN pip3 install "awscli==1.18.*" && \
    aws --version

ENV CST_VERSION 1.8.0
RUN curl -f -Lo container-structure-test https://storage.googleapis.com/container-structure-test/v${CST_VERSION}/container-structure-test-linux-amd64 && \
  chmod +x container-structure-test && \
  mv container-structure-test /usr/bin

ENV KUBECTL_VERSION 1.15.0
RUN curl -LO https://storage.googleapis.com/kubernetes-release/release/v${KUBECTL_VERSION}/bin/linux/amd64/kubectl && \
  chmod +x ./kubectl && \
  mv ./kubectl /usr/local/bin/kubectl

ENV HELM_VERSION 3.2.0
RUN curl -f -L https://get.helm.sh/helm-v${HELM_VERSION}-linux-amd64.tar.gz | tar -zx && \
  chmod +x linux-amd64/helm && \
  mv linux-amd64/helm /usr/bin && \
  rm -rf linux-amd64
