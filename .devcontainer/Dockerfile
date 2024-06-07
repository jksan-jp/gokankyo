FROM golang:1.20

# Install additional dependencies if needed
RUN apt-get update && apt-get install -y \
    git \
    && rm -rf /var/lib/apt/lists/*

# Set the working directory
WORKDIR /workspace

# Copy go.mod and run go mod tidy to generate go.sum
COPY src/go.mod ./
RUN go mod tidy

# Copy the rest of the application source code
COPY src/ .

# Expose the application port
EXPOSE 8080

# Run the application
# CMD ["go", "run", "main.go"]
