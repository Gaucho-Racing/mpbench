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

# Install git and docker client
RUN apk add --no-cache git docker-cli

WORKDIR /

COPY --from=builder /mpbench /mpbench

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
ENV TZ=America/Los_Angeles

ENTRYPOINT ["/mpbench"]