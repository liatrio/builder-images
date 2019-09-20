FROM golang:1.13-alpine AS build

RUN apk add git

RUN addgroup -S gitops && adduser -S gitops -G gitops
USER gitops

RUN go get gopkg.in/src-d/go-git.v4/... github.com/hashicorp/hcl golang.org/x/oauth2 github.com/google/go-github/github

COPY --chown=gitops:gitops ./go /go
WORKDIR /go

RUN go install gitops

FROM alpine

RUN addgroup -S gitops && adduser -S gitops -G gitops
USER gitops

COPY --from=build /go/bin /go/bin

WORKDIR /home/gitops

ENTRYPOINT [ "/go/bin/gitops" ]