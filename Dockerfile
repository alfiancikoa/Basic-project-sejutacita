# Start from golang base image
FROM golang:alpine
# Enable go modules
ENV GO111MODULE=on

# Install git. (alpine image does not have git in it)
RUN apk update && apk add --no-cache git

# Set current working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./
COPY go.sum ./

# Download all dependencies.
RUN go mod download

# Now, copy the source code
COPY . .

# Build the application.
RUN go build -o main . 
# Run executable
# CMD ["/app/main"]
CMD ./main