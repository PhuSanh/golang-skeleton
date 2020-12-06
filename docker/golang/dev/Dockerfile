FROM golang:1.15


# Set the current working directory inside the container
WORKDIR /usr/src/app/

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Install migration library
RUN go get github.com/rubenv/sql-migrate/...

# Install auto reload
RUN go get github.com/githubnemo/CompileDaemon
#RUN go get github.com/oxequa/realize

# Install swaggo
RUN go get github.com/swaggo/swag/cmd/swag

# Install mockery
RUN go get github.com/vektra/mockery/v2/.../

RUN go mod download

# Build the Go app
#RUN go build -o server
RUN go build -o ./cmd/server/server ./cmd/server/main.go