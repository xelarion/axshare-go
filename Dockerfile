# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="ErvinCheung <1390838101@qq.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# proxy
ENV GOPROXY="https://goproxy.io"

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN  go build -o axshare_go_app cmd/axshare_go.go

# Expose port 10524 to the outside world
EXPOSE 10524

# Command to run the executable
CMD ["./axshare_go_app"]