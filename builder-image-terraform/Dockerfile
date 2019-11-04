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
RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin

ENV TERRAFORM_VERSION 0.12.13
RUN curl -f -Lo terraform.zip https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip && \
  unzip terraform.zip && \
  rm -f terraform.zip && \
  chmod +x terraform && \
  mv terraform /usr/bin

ENV TERRAGRUNT_VERSION 0.21.2
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
