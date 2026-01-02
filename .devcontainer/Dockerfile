FROM mcr.microsoft.com/devcontainers/go:1.23

RUN apt-get update && apt-get install -y \
    git \
    make \
 && rm -rf /var/lib/apt/lists/*

ENV TZ=Asia/Tokyo
