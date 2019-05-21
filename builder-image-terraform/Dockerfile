FROM alpine:3.9
#FROM hashicorp/terraform:0.11.14

RUN apk add --no-cache \
    git \
    curl \
    unzip \
    openssh-client \
    make

ENV TERRAFORM_VERSION 0.10.0
RUN curl -f -Lo terraform.zip https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip && \
  unzip terraform.zip && \
  rm -f terraform.zip && \
  chmod +x terraform && \
  mv terraform /usr/bin

ENV TERRAGRUNT_VERSION 0.18.6
RUN curl -f -Lo terragrunt https://github.com/gruntwork-io/terragrunt/releases/download/v${TERRAGRUNT_VERSION}/terragrunt_linux_amd64 && \
  chmod +x terragrunt && \
  mv terragrunt /usr/bin

ENV CST_VERSION 1.8.0
RUN curl -f -Lo container-structure-test https://storage.googleapis.com/container-structure-test/v${CST_VERSION}/container-structure-test-linux-amd64 && \
  chmod +x container-structure-test && \
  mv container-structure-test /usr/bin
