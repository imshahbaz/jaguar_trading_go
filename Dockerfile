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

# Step 2: Add the Google Chrome repository and install Chrome
RUN curl -sSL https://dl.google.com/linux/linux_signing_key.pub | tee /etc/apt/trusted.gpg.d/google.asc
RUN echo "deb [arch=amd64] https://dl.google.com/linux/chrome/deb/ stable main" > /etc/apt/sources.list.d/google-chrome.list
RUN apt-get update && apt-get install -y google-chrome-stable

# Step 3: Set up VNC server and Chrome startup script
RUN mkdir -p /root/.vnc \
    && echo "password" | vncpasswd -f > /root/.vnc/passwd \
    && chmod 600 /root/.vnc/passwd

# Create a script to start the VNC server and Google Chrome
RUN echo '#!/bin/bash\n\
vncserver :1 -geometry 1280x1024 -depth 24\n\
google-chrome-stable --no-sandbox --remote-debugging-port=9222 --disable-gpu --headless\n' > /root/start.sh

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
EXPOSE 5901  # Expose VNC port

# Start VNC server and your Go application
CMD /root/start.sh && tail -f /dev/null

# Step 8: Set the command to run the application
# CMD ["./main"]
