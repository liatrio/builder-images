# Download and verify Terraform
FROM alpine:3.15 AS TERRAFORM

RUN apk --no-cache add gpgme

ENV TERRAFORM_VERSION 1.2.1
COPY sig/hashicorp.asc .
ADD https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip .
ADD https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_SHA256SUMS .
ADD https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_SHA256SUMS.sig .

RUN gpg --import hashicorp.asc \
 && gpg --verify terraform_${TERRAFORM_VERSION}_SHA256SUMS.sig \
 && cat terraform_${TERRAFORM_VERSION}_SHA256SUMS | grep linux_amd64 | sha256sum -c \
 && unzip terraform_${TERRAFORM_VERSION}_linux_amd64.zip \
 && chmod +x terraform

# Download and verify AWS IAM Authenticator
FROM alpine:3.15 as AWS_IAM_AUTHENTICATOR

RUN apk add --no-cache openssl

ENV AWS_IAM_AUTHENTICATOR_VERSION 1.21.2/2021-07-05
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
ENV KUBECTL_VERSION v1.24.0
ADD https://storage.googleapis.com/kubernetes-release/release/${KUBECTL_VERSION}/bin/linux/amd64/kubectl /usr/bin
RUN chmod a+x /usr/bin/kubectl

## Install AWS IAM Authenticator
COPY --from=AWS_IAM_AUTHENTICATOR /aws-iam-authenticator /usr/bin/

## Create Jenkins user
RUN addgroup -g 1000 jenkins && adduser -h /home/jenkins -G jenkins -u 1000 -D jenkins
USER jenkins
WORKDIR /home/jenkins
