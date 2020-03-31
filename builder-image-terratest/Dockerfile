# Download and verify Terraform
FROM alpine:3.11 AS TERRAFORM

RUN apk --no-cache add gpgme

ENV TERRAFORM_VERSION 0.12.24
COPY sig/hashicorp.asc .
ADD https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip .
ADD https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_SHA256SUMS .
ADD https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_SHA256SUMS.sig .

RUN gpg --import hashicorp.asc \
 && gpg --verify terraform_${TERRAFORM_VERSION}_SHA256SUMS.sig \
 && cat terraform_${TERRAFORM_VERSION}_SHA256SUMS | grep linux_amd64 | sha256sum -c \
 && unzip terraform_${TERRAFORM_VERSION}_linux_amd64.zip \
 && chmod +x terraform

# Download Keycloak && Harbor Terraform Providers
FROM alpine:3.11 AS TFPROVIDERS

WORKDIR /plugins

ENV TERRAFORM_PROVIDER_KEYCLOAK_VERSION 1.11.1
ADD https://github.com/mrparkers/terraform-provider-keycloak/releases/download/${TERRAFORM_PROVIDER_KEYCLOAK_VERSION}/terraform-provider-keycloak_v${TERRAFORM_PROVIDER_KEYCLOAK_VERSION}_linux_amd64_static.zip /
RUN unzip -j /terraform-provider-keycloak_v${TERRAFORM_PROVIDER_KEYCLOAK_VERSION}_linux_amd64_static.zip

ENV TERRAFORM_PROVIDER_HARBOR_VERSION 0.0.4
ADD https://github.com/liatrio/terraform-provider-harbor/releases/download/v${TERRAFORM_PROVIDER_HARBOR_VERSION}/terraform-provider-harbor_${TERRAFORM_PROVIDER_HARBOR_VERSION}_linux_amd64.zip /
RUN unzip -j /terraform-provider-harbor_${TERRAFORM_PROVIDER_HARBOR_VERSION}_linux_amd64.zip -x "LICENSE"

# Download and verify AWS IAM Authenticator
FROM alpine:3.11 as AWS_IAM_AUTHENTICATOR

RUN apk add --no-cache openssl

ENV AWS_IAM_AUTHENTICATOR_VERSION 1.14.6/2019-08-22
ADD https://amazon-eks.s3-us-west-2.amazonaws.com/${AWS_IAM_AUTHENTICATOR_VERSION}/bin/linux/amd64/aws-iam-authenticator /
ADD https://amazon-eks.s3-us-west-2.amazonaws.com/${AWS_IAM_AUTHENTICATOR_VERSION}/bin/linux/amd64/aws-iam-authenticator.sha256 /

RUN openssl sha1 -sha256 /aws-iam-authenticator

RUN chmod +x /aws-iam-authenticator

# Main image includes golang, terraform, kubectl, Keycloak
FROM alpine:3.11

RUN apk add --no-cache \
    curl \
    git \
    make \
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

## Install AWS IAM Authenticator
COPY --from=AWS_IAM_AUTHENTICATOR /aws-iam-authenticator /usr/bin/

## Create Jenkins user
RUN addgroup -g 1000 jenkins && adduser -h /home/jenkins -G jenkins -u 1000 -D jenkins
USER jenkins
WORKDIR /home/jenkins

## Install Terraform plugins
COPY --chown=jenkins:jenkins --from=TFPROVIDERS /plugins/* /home/jenkins/.terraform.d/plugins/
