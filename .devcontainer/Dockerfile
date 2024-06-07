FROM golang:1.20

# Install additional dependencies if needed
RUN apt-get update && apt-get install -y \
    git \
    && rm -rf /var/lib/apt/lists/*

# Set the working directory
WORKDIR /workspaceFROM golang:1.20

# Install additional dependencies if needed
RUN apt-get update && apt-get install -y \
    git \
    && rm -rf /var/lib/apt/lists/*

# Set the working directory
WORKDIR /workspace

# Copy go.mod and go.sum files
COPY src/go.mod src/go.sum ./

# Download dependencies
RUN go mod tidy

# Copy the rest of the application source code
COPY src/ .

# Expose the application port
EXPOSE 8080

# Run the application
RUN ["go", "run", "main.go"]

# Install Go MySQL driver
RUN go get -u github.com/go-sql-driver/mysql

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["go", "run", "main.go"]
