FROM ubuntu:22.04

LABEL org.opencontainers.image.source https://github.com/liatrio/builder-images

ARG GO_VERSION=1.18.2
ARG YQ_VERSION=4.5.1
ARG TFLINT_VERSION=v0.36.2
ARG TERRAFORM_VERSION=1.2.1
ARG TERRAGRUNT_VERSION=v0.37.1
 
ENV PATH="$PATH:/usr/local/go/bin"

RUN apt-get update && apt-get install -yq \
    apt-transport-https \
    ca-certificates \
    curl \
    git \
    gnupg \
    jq \
    lsb-release \
    wget \
    zip && \
    rm -rf /var/lib/apt/lists/*

# Install yq
RUN cd /tmp && \
    wget https://github.com/mikefarah/yq/releases/download/v${YQ_VERSION}/yq_linux_amd64 && \
    chmod +x yq_linux_amd64 && \
    mv -vf yq_linux_amd64 /usr/local/bin/yq && \
    yq -V

# Install azure cli
RUN curl -sL https://packages.microsoft.com/keys/microsoft.asc | \
    gpg --dearmor | \
    tee /etc/apt/trusted.gpg.d/microsoft.gpg > /dev/null && \
    echo "deb [arch=amd64] https://packages.microsoft.com/repos/azure-cli/ $(lsb_release -cs) main" | \
    tee /etc/apt/sources.list.d/azure-cli.list && \
    apt-get update && apt-get install -yq azure-cli && \
    rm -rf /var/lib/apt/lists/* && \
    az extension add --system --name azure-devops && \
    az -v

# Install go
RUN wget -O /tmp/go.tar.gz https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf /tmp/go.tar.gz && \
    rm /tmp/go.tar.gz && \
    go version

# Install tfenv and setup gpg verification
RUN mkdir /opt/tfenv && \
    git clone https://github.com/tfutils/tfenv.git /opt/tfenv/.tfenv && \
    chmod -R ago+w /opt/tfenv/ && \ 
    echo 'PATH="/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games:/opt/tfenv/.tfenv/bin"' > /etc/environment && \
    touch /opt/tfenv/.tfenv/use-gnupg && \
    curl https://keybase.io/hashicorp/pgp_keys.asc | gpg --import && \
    echo 'curl https://keybase.io/hashicorp/pgp_keys.asc | gpg --import' >> /etc/skel/.profile

ENV PATH $PATH:/usr/local/go/bin:/opt/tfenv/.tfenv/bin
 
# Install terraform
RUN tfenv install ${TERRAFORM_VERSION} && \
    tfenv use ${TERRAFORM_VERSION} && \
    chmod -R ago+w /opt/tfenv/ && \
    terraform version

# Install terragrunt
RUN wget -O /tmp/terragrunt_linux_amd64 https://github.com/gruntwork-io/terragrunt/releases/download/${TERRAGRUNT_VERSION}/terragrunt_linux_amd64 && \
    mv /tmp/terragrunt_linux_amd64 /usr/bin/terragrunt && \
    chmod +x /usr/bin/terragrunt && \
    terragrunt --version

# Install tflint
RUN mkdir -p /opt/tflint && \
    cd /opt/tflint && \
    wget -O tflint.zip https://github.com/terraform-linters/tflint/releases/download/${TFLINT_VERSION}/tflint_linux_amd64.zip && \
    unzip tflint.zip && \
    chmod +x tflint && \
    mv -vf tflint /usr/local/bin/tflint && \
    rm -rf /opt/tflint && \
    tflint -v
