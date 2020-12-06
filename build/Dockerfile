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
  go build -mod=vendor -o cli.bin cmd/main.go

# Second step to build minimal image
FROM alpine:3.12
RUN apk add --update --no-cache ca-certificates

RUN apk --update add supervisor && \
  mkdir -p /var/log/supervisor && \
  mkdir -p /etc/supervisor/conf.d

COPY ./debug/config/ /etc/

RUN mkdir /app
WORKDIR /app
COPY --from=builder /build/cli.bin .

# default command 
CMD ["supervisord", "-n", "-c", "/etc/supervisor/supervisord.conf"]