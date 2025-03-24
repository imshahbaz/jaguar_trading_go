# # Step 1: Build the Go application in a Go-based container
# FROM golang:1.24.1-alpine AS builder

# COPY . /app
# # Step 2: Set the working directory
# WORKDIR /app

# # Step 3: Copy the Go modules and download dependencies
# COPY go.mod go.sum ./
# RUN go mod download

# # Step 4: Copy the source code into the container
# COPY . .

# # Step 5: Build the Go application (the output is a binary file)
# RUN go build -o main .

# # Step 6: Create a smaller image for the runtime
# FROM alpine:latest

# FROM python:3.9-alpine AS runtime

# # Step 7: Set the working directory in the runtime container
# WORKDIR /root/

# # Step 8: Copy the compiled binary from the builder container
# COPY --from=builder /app/main .

# # Step 9: Expose the port the app will run on
# EXPOSE 8080

# # Step 10: Set the command to run the application
# CMD ["./main"]





# Step 1: Build the Go application in a Go-based container
FROM golang:1.24.1-alpine AS builder

# Step 2: Set the working directory
WORKDIR /app

# Step 3: Copy go.mod and go.sum, and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Step 4: Copy the source code into the container
COPY . .

# Step 5: Build the Go application (the output is a binary file)
RUN go build -o main .

# Step 6: Create a smaller image for the runtime (use python:3.9-alpine or python2-alpine)
FROM python:3.9-alpine AS runtime

# Step 7: Set the working directory in the runtime container
WORKDIR /root/

# Step 8: Copy the compiled binary from the builder container
COPY --from=builder /app/main .

# Step 9: Copy the Python script and requirements.txt into the container
COPY ./get_data.py ./get_data.py
COPY ./requirements.txt ./requirements.txt  # Copy the requirements.txt into the container

# Step 10: Install system dependencies for building Python packages (e.g., for cffi)
RUN apk add --no-cache --virtual .build-deps gcc musl-dev libffi-dev

# Step 11: Install Python dependencies from requirements.txt
RUN pip install --no-cache-dir -r requirements.txt

# Step 12: Clean up build dependencies (optional but recommended)
RUN apk del .build-deps

# Step 13: Create a symlink for python to point to python3 (if needed)
RUN ln -sf python3 /usr/bin/python

# Step 14: Expose the port the app will run on
EXPOSE 8080

# Step 15: Set the command to run the Go application (Fiber app)
CMD ["./main"]
