#FROM alpine:3.9
FROM hashicorp/terraform:0.11.14

RUN apk add --no-cache curl

RUN curl -f -Lo terragrunt https://github.com/gruntwork-io/terragrunt/releases/download/v0.18.6/terragrunt_linux_amd64 && \
  chmod +x terragrunt && \
  mv terragrunt /usr/bin

ENV CST_VERSION 1.8.0
RUN curl -f -Lo container-structure-test https://storage.googleapis.com/container-structure-test/v${CST_VERSION}/container-structure-test-linux-amd64 && \
  chmod +x container-structure-test && \
  mv container-structure-test /usr/bin
