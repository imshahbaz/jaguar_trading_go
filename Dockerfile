FROM ubuntu:latest

# Step 1: Install dependencies for Chrome and VNC
RUN apt-get update && apt-get install -y \
    wget \
    curl \
    gnupg2 \
    ca-certificates \
    fonts-liberation \
    libappindicator3-1 \
    libnspr4 \
    libnss3 \
    libxss1 \
    libxtst6 \
    xdg-utils \
    x11vnc \
    xfce4 \
    tightvncserver \
    && rm -rf /var/lib/apt/lists/*

# Step 2: Install Google Chrome
RUN wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb \
    && dpkg -i google-chrome-stable_current_amd64.deb \
    && apt-get install -f -y \
    && rm google-chrome-stable_current_amd64.deb

# Step 3: Set up VNC server and Chrome startup script
RUN mkdir -p /root/.vnc \
    && echo "password" | vncpasswd -f > /root/.vnc/passwd \
    && chmod 600 /root/.vnc/passwd

# Create a script to start the VNC server and Google Chrome
RUN echo '#!/bin/bash\n\
vncserver :1 -geometry 1280x1024 -depth 24\n\
/opt/google/chrome/google-chrome-stable --no-sandbox --remote-debugging-port=9222 --disable-gpu --headless\n' > /root/start.sh

RUN chmod +x /root/start.sh

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

# Step 8: Set the command to run the application
CMD ["./main"]
