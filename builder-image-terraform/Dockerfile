FROM alpine:3.9 AS buildKeycloak

RUN apk add \
    git \
    go \
    libc-dev \
    make

RUN addgroup -g 1000 jenkins && adduser -h /home/jenkins -G jenkins -u 1000 -D jenkins
USER jenkins
WORKDIR /home/jenkins

RUN wget https://github.com/mrparkers/terraform-provider-keycloak/archive/master.zip -O - | unzip -

# the 2nd line is a temporary fix because git.apache.org is gone. see https://github.com/mrparkers/terraform-provider-keycloak/pull/150
RUN	cd ./terraform-provider-keycloak-master \
    && echo "replace git.apache.org/thrift.git => github.com/apache/thrift v0.0.0-20180902110319-2566ecd5d999" >> go.mod \
    && GO111MODULE=on go mod download \
    && make build

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

ENV TERRAFORM_VERSION 0.12.3
RUN curl -f -Lo terraform.zip https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip && \
  unzip terraform.zip && \
  rm -f terraform.zip && \
  chmod +x terraform && \
  mv terraform /usr/bin

ENV TERRAGRUNT_VERSION 0.19.8
RUN curl -f -Lo terragrunt https://github.com/gruntwork-io/terragrunt/releases/download/v${TERRAGRUNT_VERSION}/terragrunt_linux_amd64 && \
  chmod +x terragrunt && \
  mv terragrunt /usr/bin

RUN addgroup -g 1000 jenkins && adduser -h /home/jenkins -G jenkins -u 1000 -D jenkins
USER jenkins

COPY --from=buildKeycloak --chown=1000:1000 /home/jenkins/terraform-provider-keycloak-master/terraform-provider-keycloak /home/jenkins/.terraform.d/plugins/
