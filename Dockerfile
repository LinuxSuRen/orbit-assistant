FROM golang:1.17 as builder

WORKDIR /workspace
COPY . .
RUN go mod download
RUN CGO_ENABLE=0 go build -ldflags "-w -s" -o orbit

FROM alpine:3.10

LABEL "com.github.actions.name"="Orbit Assistant"
LABEL "com.github.actions.description"="Orbit Assistant"
LABEL "com.github.actions.icon"="home"
LABEL "com.github.actions.color"="red"

LABEL "repository"="https://github.com/linuxsuren/orbit-assistant"
LABEL "homepage"="https://github.com/linuxsuren/orbit-assistant"
LABEL "maintainer"="Rick <linuxsuren@gmail.com>"

LABEL "Name"="Orbit Assistant"
LABEL "Version"="0.0.1"

ENV LC_ALL C.UTF-8
ENV LANG en_US.UTF-8
ENV LANGUAGE en_US.UTF-8

RUN apk add --no-cache \
        git \
        openssh-client \
        libc6-compat \
        libstdc++

COPY entrypoint.sh /entrypoint.sh
COPY --from=builder /workspace/orbit /usr/bin/orbit

ENTRYPOINT ["/entrypoint.sh"]
