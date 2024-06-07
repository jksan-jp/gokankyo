FROM golang:1.20

# Install additional dependencies if needed
RUN apt-get update && apt-get install -y \
    git \
    && rm -rf /var/lib/apt/lists/*

# Set the working directory
WORKDIR /workspace

# Copy the rest of the application source code
COPY src/* ./
RUN go mod download

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["go", "run", "main.go"]
