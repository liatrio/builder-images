FROM summerwind/actions-runner:v2.291.1-ubuntu-20.04-3ca1152

LABEL org.opencontainers.image.source https://github.com/liatrio/builder-images

USER root
WORKDIR /usr/workspace

# AWS CLI is needed to authenticate to EKS clusters, and to assume a role in CI
ENV AWS_PAGER=""
ENV AWS_REGION="us-east-1"
RUN curl -LO https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip && \
    unzip awscli-exe-linux-x86_64.zip && \
    ./aws/install && \
    rm -rf aws awscli-exe-linux-x86_64.zip && \
    aws --version

# Install azure cli
RUN curl -sL https://packages.microsoft.com/keys/microsoft.asc | \
    gpg --dearmor | \
    tee /etc/apt/trusted.gpg.d/microsoft.gpg > /dev/null && \
    echo "deb [arch=amd64] https://packages.microsoft.com/repos/azure-cli/ $(lsb_release -cs) main" | \
    tee /etc/apt/sources.list.d/azure-cli.list && \
    apt-get update && apt-get install -yq azure-cli && \
    az extension add --system --name azure-devops && \
    az -v && \
    chown runner -R /home/runner

# Go is needed to run terraform tests via terratest
ENV PATH="${PATH}:/usr/local/go/bin"
ENV GO_VERSION="1.18.2"
RUN curl -LO https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz && \
    rm go${GO_VERSION}.linux-amd64.tar.gz && \
    go version

ENV TERRAFORM_VERSION="1.2.1"
RUN curl -LO https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip && \
    unzip terraform_${TERRAFORM_VERSION}_linux_amd64.zip && \
    mv ./terraform /usr/local/bin && \
    rm terraform_${TERRAFORM_VERSION}_linux_amd64.zip && \
    terraform -version

ENV TERRAGRUNT_VERSION="0.37.1"
RUN curl -LO https://github.com/gruntwork-io/terragrunt/releases/download/v${TERRAGRUNT_VERSION}/terragrunt_linux_amd64 && \
    chmod +x ./terragrunt_linux_amd64 && \
    mv ./terragrunt_linux_amd64 /usr/local/bin/terragrunt && \
    terragrunt --version

ENV VCLUSTER_VERSION="0.8.1"
RUN curl -LO https://github.com/loft-sh/vcluster/releases/download/v${VCLUSTER_VERSION}/vcluster-linux-amd64 && \
    chmod +x ./vcluster-linux-amd64 && \
    mv ./vcluster-linux-amd64 /usr/local/bin/vcluster && \
    vcluster --version

ENV KUBECTL_VERSION="1.24.0"
RUN curl -LO https://dl.k8s.io/release/v${KUBECTL_VERSION}/bin/linux/amd64/kubectl && \
    chmod +x ./kubectl && \
    mv ./kubectl /usr/local/bin && \
    kubectl version --client

USER runner
