FROM golang:alpine AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux

#RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN apk update && apk upgrade && apk add --no-cache ca-certificates
RUN update-ca-certificates

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build  -o . cmd/writeMetric/main.go

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/main .

# Build a small image
FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /dist/main /

# Command to run
ENTRYPOINT ["/main"]
