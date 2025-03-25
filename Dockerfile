# Step 1: Build the Go application in a Go-based container
FROM golang:1.24.1-alpine AS builder

# Set the working directory
WORKDIR /app

# Step 2: Copy the Go modules files and download dependencies
COPY go.mod go.sum ./
RUN go mod tidy && go mod download

# Step 3: Copy the rest of the source code
COPY . .

# Step 4: Build the Go application (the output is a binary file)
RUN go build -o main .

# Step 5: Create a smaller image for the runtime
FROM alpine:latest

# Set the working directory in the runtime container
WORKDIR /root/

# Step 6: Copy the compiled binary from the builder container
COPY --from=builder /app/main .

# Step 7: Expose the port the app will run on
EXPOSE 8080

Step 8: Set the command to run the application
CMD ["./main"]
