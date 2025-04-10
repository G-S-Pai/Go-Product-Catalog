# Use an official Golang runtime as a parent image
FROM golang:1.24

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules and download the dependencies
COPY go.mod ./
RUN go mod download

# Copy the entire project into the container
COPY . .

# Build the Go app
RUN go build -o github.com/g-s-pai/go-product-catalog .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./github.com/g-s-pai/go-product-catalog"]
