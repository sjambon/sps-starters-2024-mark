# Use the official Golang image to create a build artifact.
FROM golang:1.16 AS builder

# Set the working directory outside $GOPATH to enable the support for modules.
WORKDIR /src

# Copy go mod file
COPY go.mod ./

# Copy go sum file if it exists
COPY go.sum* ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch
FROM alpine:latest  

RUN apk --no-cache add ca-certificates postgresql-client

WORKDIR /app/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /src/main .
COPY wait-for-it.sh .
RUN chmod +x wait-for-it.sh

# Expose port 8000 to the outside world
EXPOSE 8000

# Command to run the executable
CMD ["./main"]