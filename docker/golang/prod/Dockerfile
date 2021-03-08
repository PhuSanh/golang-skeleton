# Start from golang base image
FROM golang:1.15-alpine as builder

# ENV GO111MODULE=on

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
RUN apk add build-base

# Set the current working directory inside the container
#WORKDIR /app
WORKDIR /usr/src/app/

# Copy go mod and sum files
COPY go.mod go.sum ./

# RUN go mod tidy

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

# Install migration library
RUN go get github.com/rubenv/sql-migrate/...

# Install swaggo
#RUN go get github.com/swaggo/swag/cmd/swag

# Install mockery
#RUN go get github.com/vektra/mockery/v2/.../

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./cmd/server/server ./cmd/server/main.go

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN apk add bash

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /usr/src/app/cmd/server/server .
COPY --from=builder /usr/src/app/.env .
COPY --from=builder /usr/src/app/tools/scripts ./tools/scripts
COPY --from=builder /usr/src/app/databases/migrations ./databases/migrations
COPY --from=builder /go/bin/sql-migrate ./bin/sql-migrate

# export root/bin to execute file in bin folder
ENV PATH "$PATH:/root/bin"