# Download and verify Terraform
FROM alpine:3.11 AS TERRAFORM

RUN apk --update add gpgme

ENV TERRAFORM_VERSION 0.12.20
COPY sig/hashicorp.asc .
ADD https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip .
ADD https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_SHA256SUMS .
ADD https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_SHA256SUMS.sig .

RUN gpg --import hashicorp.asc \
 && gpg --verify terraform_${TERRAFORM_VERSION}_SHA256SUMS.sig \
 && cat terraform_${TERRAFORM_VERSION}_SHA256SUMS | grep linux_amd64 | sha256sum -c \
 && unzip terraform_${TERRAFORM_VERSION}_linux_amd64.zip \
 && chmod +x terraform

# Download Keycloak Terraform Provider
FROM alpine:3.11 AS KEYCLOAK

WORKDIR /plugins

ENV TERRAFORM_PROVIDER_KEYCLOAK_VERSION 1.11.1
ADD https://github.com/mrparkers/terraform-provider-keycloak/releases/download/${TERRAFORM_PROVIDER_KEYCLOAK_VERSION}/terraform-provider-keycloak_v${TERRAFORM_PROVIDER_KEYCLOAK_VERSION}_linux_amd64_static.zip /
RUN unzip -j /terraform-provider-keycloak_v${TERRAFORM_PROVIDER_KEYCLOAK_VERSION}_linux_amd64_static.zip

# Main image includes golang, terraform, kubectl, Keycloak
FROM alpine:3.11

RUN apk add --no-cache \
    go

## Configure Go
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH
RUN mkdir -m 0777 ${GOPATH} ${GOPATH}/src ${GOPATH}/bin

## Install Terraform
COPY --from=TERRAFORM terraform /usr/bin

## Install Kubectl
ENV KUBECTL_VERSION v1.17.0
ADD https://storage.googleapis.com/kubernetes-release/release/${KUBECTL_VERSION}/bin/linux/amd64/kubectl /usr/bin
RUN chmod a+x /usr/bin/kubectl

## Create Jenkins user
RUN addgroup -g 1000 jenkins && adduser -h /home/jenkins -G jenkins -u 1000 -D jenkins
USER jenkins
WORKDIR /home/jenkins

## Install Keycloak plugin
COPY --chown=jenkins:jenkins --from=KEYCLOAK /plugins/* /home/jenkins/.terraform.d/plugins/
