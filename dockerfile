FROM golang:1.19-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module and sum files to the container
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of your application source code to the container
COPY . /app

# Build your Go application
RUN go build -o app

# Expose the port your application will listen on
EXPOSE 8000

# Define the command to run your application
CMD ["./app"]
