# Step 1: Build the Go application in a Go-based container
FROM golang:1.24.1-alpine AS builder

COPY . /app
# Step 2: Set the working directory
WORKDIR /app

# Step 3: Copy the Go modules and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Step 4: Copy the source code into the container
COPY . .

# Step 5: Build the Go application (the output is a binary file)
RUN go build -o main .

# Step 6: Create a smaller image for the runtime
FROM alpine:latest

# Step 7: Set the working directory in the runtime container
WORKDIR /root/

# Step 8: Copy the compiled binary from the builder container
COPY --from=builder /app/main .

# Step 9: Expose the port the app will run on
EXPOSE 8080

# Step 10: Set the command to run the application
CMD ["./main"]