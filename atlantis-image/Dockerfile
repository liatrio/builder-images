FROM ghcr.io/runatlantis/atlantis:v0.19.3

RUN apk add --no-cache aws-cli

ENV TERRAGRUNT_VERSION 0.37.1
RUN curl -f -Lo terragrunt https://github.com/gruntwork-io/terragrunt/releases/download/v${TERRAGRUNT_VERSION}/terragrunt_linux_amd64 && \
  chmod +x terragrunt && \
  mv terragrunt /usr/bin && \
  terragrunt -version
