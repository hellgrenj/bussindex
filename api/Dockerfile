FROM golang:1.15-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main ./cmd/webapi/main.go

# Expose port 8080 to the outside world
EXPOSE 8080

FROM alpine:3.11 as runtime

# copy only executable from build step
COPY --from=builder ./app/main main

# Command to run the executable
CMD ["./main"]
