# Alpine
FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy the source code
COPY . ./

# Build 
RUN go build -o auth-service cmd/main.go

# Start service
CMD ["./auth-service"]
