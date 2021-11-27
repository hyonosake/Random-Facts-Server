# Start from golang base image
FROM golang:alpine as builder

WORKDIR go/src/github.com/APIserver

# Enable go modules
ENV GO111MODULE=on

# Install git. (alpine image does not have git in it)
RUN apk update && apk add --no-cache git

# Copy go mod and sum files
COPY go.mod .
COPY go.sum .

# Download all dependencies.
RUN go mod download

# Now, copy the source code
COPY . .

# Note here: CGO_ENABLED is disabled for cross system compilation
# It is also a common best practise.

# Build the application.
RUN cd server && go build -o server .

CMD ["./server/server"]
