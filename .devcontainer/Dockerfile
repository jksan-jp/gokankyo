FROM golang:1.20

# Install additional dependencies if needed
RUN apt-get update && apt-get install -y \
    git \
    && rm -rf /var/lib/apt/lists/*

# Set the working directory
WORKDIR /workspace

# Install Go MySQL driver
RUN go get -u github.com/go-sql-driver/mysql

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["go", "run", "main.go"]
