FROM golang:1.24.4-alpine3.21

WORKDIR /src/app

# Install Air
RUN go install github.com/air-verse/air@latest

# Copy go mod and sum files first for better caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Ensure go.mod is up to date
RUN go mod tidy

EXPOSE 8080