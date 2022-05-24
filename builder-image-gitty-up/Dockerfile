FROM alpine:3.15

LABEL org.opencontainers.image.source https://github.com/liatrio/builder-images

ENV GITTYUP_VERSION 0.1.5
RUN wget https://github.com/liatrio/gitty-up/releases/download/v${GITTYUP_VERSION}/gitty-up_${GITTYUP_VERSION}_Linux_x86_64.tar.gz -O - | tar -z -x gitty-up -C /usr/local/bin/

RUN addgroup -g 1000 gitops && adduser -h /home/gitops -G gitops -u 1000 -D gitops
USER gitops
WORKDIR /home/gitops

ENTRYPOINT [ "/usr/local/bin/gitops" ]
