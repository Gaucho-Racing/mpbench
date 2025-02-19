FROM --platform=$BUILDPLATFORM golang:1.23-alpine AS builder

RUN apk --no-cache add ca-certificates
RUN apk add --no-cache tzdata

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
ARG TARGETOS
ARG TARGETARCH
RUN GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o /mpbench

##
## Deploy
##
FROM alpine:3.21

# Install git and newer docker client from the community repository
RUN apk add --no-cache git
RUN apk add --no-cache docker

# Script to handle docker login
COPY docker-login.sh /docker-login.sh
RUN chmod +x /docker-login.sh

WORKDIR /

COPY --from=builder /mpbench /mpbench

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
ENV TZ=America/Los_Angeles

# Use shell form to allow environment variable expansion
ENTRYPOINT ["/bin/sh", "-c", "/docker-login.sh && /mpbench"]