# Base image
FROM golang:latest

# Konteyner ichida ishchi katalogini o'rnating
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main ./api/cmd/main.go

# Command to run the executable
CMD ["./main"]
