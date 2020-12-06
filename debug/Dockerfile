FROM golang:1.12.1-stretch as builder

# All these steps will be cached
WORKDIR /build
# COPY go.mod and go.sum files to the workspace
COPY go.mod go.sum ./
# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download
# COPY the source code as the last step
COPY . .
# This step is required to create dir vendor
RUN go mod vendor
# Build the binary
RUN CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64 \
  go build -mod=vendor -gcflags="all=-N -l" -o cli.bin cmd/main.go

# Second step to build minimal image
FROM alpine:3.12
RUN apk add --no-cache ca-certificates \
  dpkg \
  gcc \
  git \
  musl-dev \
  bash \
  make \
  libc6-compat \
  git

RUN wget https://dl.google.com/go/go1.14.4.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.14.4.linux-amd64.tar.gz
RUN rm go1.14.4.linux-amd64.tar.gz
RUN export PATH=$PATH:/usr/local/go/bin

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
ENV GO111MODULE on

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" \
  && chmod -R 777 "$GOPATH"

RUN go version

RUN go get github.com/go-delve/delve/cmd/dlv

RUN apk --update add supervisor && \
  mkdir -p /var/log/supervisor && \
  mkdir -p /etc/supervisor/conf.d

COPY ./debug/config/ /etc/

RUN mkdir /app
WORKDIR /app
COPY --from=builder /build/cli.bin .

# default command 
CMD ["supervisord", "-n", "-c", "/etc/supervisor/supervisord.conf"]