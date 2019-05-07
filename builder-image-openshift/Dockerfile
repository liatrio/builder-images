FROM alpine:3.9

RUN apk add --no-cache libc6-compat

RUN wget https://github.com/openshift/origin/releases/download/v3.11.0/openshift-origin-client-tools-v3.11.0-0cbc58b-linux-64bit.tar.gz -O - | tar -xz
ENV PATH ${PATH}:/openshift-origin-client-tools-v3.11.0-0cbc58b-linux-64bit
