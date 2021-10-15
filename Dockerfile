FROM golang:alpine AS builder

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Install docker & docker-compose
RUN apk update
RUN apk add docker docker-compose

RUN adduser -SDHs /sbin/nologin dockremap
RUN addgroup -S dockremap
RUN echo dockremap:$(cat /etc/passwd|grep dockremap|cut -d: -f3):65536 >> /etc/subuid
RUN echo dockremap:$(cat /etc/passwd|grep dockremap|cut -d: -f4):65536 >> /etc/subgid
RUN ls -ali /etc/docker
RUN cat /etc/docker/daemon.json

# Start docker daemon
RUN dockerd
#service docker start

#docker start

#RUN ln -s /usr/local/bin/docker-compose /compose/docker-compose
RUN docker-compose -f ./config/database/docker-compose.yaml up

# Build the application
RUN go build -o main ./cmd/gpsDataCollector

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/main .

# Run tests
CMD CGO_ENABLED=0 go test ./...

# Build a small image
FROM scratch

COPY --from=builder /dist/main /

# Command to run
ENTRYPOINT ["/main"]