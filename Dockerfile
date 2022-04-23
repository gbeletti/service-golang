# --- Stage 1:
FROM golang:1.18-alpine as builder
# Args & ENVs
ENV BUILD_PATH=/go/src/github.com/gbeletti/service-golang

RUN apk update && apk add --no-cache curl gcc git libc-dev
# COPY local files
WORKDIR ${BUILD_PATH}
COPY . .

# Get go dependencies
RUN go mod download

# revive (go lint successor)
RUN go install github.com/mgechev/revive@latest && \
    revive ./...

# gosec - Golang Security Checker
RUN curl -sfL https://raw.githubusercontent.com/securego/gosec/master/install.sh | sh -s -- -b ${GOPATH}/bin latest && \
    gosec ./...

# Build dynamically linked Go binary
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 \
    go build -o service

RUN cp ${BUILD_PATH}/service /bin/service

# --- Stage 2:
FROM alpine:3
# Install dependencies
RUN apk update && apk add --no-cache ca-certificates tzdata libc6-compat
# Copy binary from builder
COPY --from=builder /bin/service /service
# Run the application on container startup.
CMD ["/service"]
EXPOSE 8000