FROM golang:1.13-alpine AS build

RUN apk add git

RUN addgroup -g 1000 jenkins && adduser -h /home/jenkins -G jenkins -u 1000 -D jenkins
USER jenkins

RUN go get gopkg.in/src-d/go-git.v4/... github.com/hashicorp/hcl golang.org/x/oauth2 github.com/google/go-github/github

COPY --chown=jenkins:jenkins ./go /go
WORKDIR /go

RUN go install gitops

FROM alpine

RUN addgroup -g 1000 jenkins && adduser -h /home/jenkins -G jenkins -u 1000 -D jenkins
USER jenkins

COPY --from=build /go/bin /go/bin

WORKDIR /home/jenkins

# ENTRYPOINT [ "/go/bin/gitops" ]