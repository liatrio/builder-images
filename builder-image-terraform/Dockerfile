FROM alpine:3.9

RUN apk add --no-cache \
    git \
    curl \
    unzip \
    openssh-client \
    make \
    musl-dev \
    go

# Configure Go
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH
RUN mkdir -m 0777 ${GOPATH} ${GOPATH}/src ${GOPATH}/bin

WORKDIR /tmp/
ENV TERRAFORM_VERSION 0.12.20
COPY sig/hashicorp.asc /tmp/.
ADD https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip /tmp/
ADD https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_SHA256SUMS /tmp/
ADD https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_SHA256SUMS.sig /tmp/

RUN apk --update add --virtual verify gpgme \
 && gpg --import hashicorp.asc \
 && gpg --verify /tmp/terraform_${TERRAFORM_VERSION}_SHA256SUMS.sig \
 && cat terraform_${TERRAFORM_VERSION}_SHA256SUMS | grep linux_amd64 | sha256sum -c \
 && unzip terraform_${TERRAFORM_VERSION}_linux_amd64.zip \
 && chmod +x terraform \
 && mv terraform /usr/bin \
 && apk del verify \
 && rm -rf /tmp/* \
 && rm -rf /var/cache/apk/*

ENV TERRAGRUNT_VERSION 0.21.10
RUN curl -f -Lo terragrunt https://github.com/gruntwork-io/terragrunt/releases/download/v${TERRAGRUNT_VERSION}/terragrunt_linux_amd64 && \
  chmod +x terragrunt && \
  mv terragrunt /usr/bin

RUN addgroup -g 1000 jenkins && adduser -h /home/jenkins -G jenkins -u 1000 -D jenkins
USER jenkins
WORKDIR /home/jenkins

ENV TERRAFORM_PROVIDER_KEYCLOAK_VERSION 1.11.1
RUN curl -Lo terraform-provider-keycloak.zip https://github.com/mrparkers/terraform-provider-keycloak/releases/download/${TERRAFORM_PROVIDER_KEYCLOAK_VERSION}/terraform-provider-keycloak_v${TERRAFORM_PROVIDER_KEYCLOAK_VERSION}_linux_amd64_static.zip && \
    unzip -j terraform-provider-keycloak.zip && \
    rm -f terraform-provider-keycloak.zip && \
    mkdir -p /home/jenkins/.terraform.d/plugins/ && \
    mv terraform-provider-keycloak_v${TERRAFORM_PROVIDER_KEYCLOAK_VERSION} /home/jenkins/.terraform.d/plugins/
