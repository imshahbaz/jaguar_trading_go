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

RUN apk update && apk add --no-cache \
    bash \
    ca-certificates \
    wget \
    curl \
    libx11 \
    libx264 \
    libfontconfig \
    libxss \
    libnss3 \
    libatk-bridge2.0-0 \
    libgtk-3-0 \
    --repository=http://dl-cdn.alpinelinux.org/alpine/edge/main

# Step 8: Install Google Chrome
RUN wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb -O /tmp/google-chrome.deb && \
    apk add --no-cache --allow-untrusted /tmp/google-chrome.deb && \
    rm /tmp/google-chrome.deb

# Step 7: Set the working directory in the runtime container
WORKDIR /root/

# Step 8: Copy the compiled binary from the builder container
COPY --from=builder /app/main .

ENV CHROME_BIN=/usr/bin/google-chrome-stable

# Step 9: Expose the port the app will run on
EXPOSE 8080

# Step 10: Set the command to run the application
CMD ["./main"]
