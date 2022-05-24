FROM maven:3.8.5-jdk-8

LABEL org.opencontainers.image.source https://github.com/liatrio/builder-images

RUN apt-get update && \
    apt-get install -y \
		git \
		curl \
		make && \
    rm -rf /var/lib/apt/lists/*

