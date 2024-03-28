# Use Golang Version
FROM golang:latest as builder

# Working Directory Setting
WORKDIR /server

# Base File Copy
COPY go.mod go.sum ./
RUN go mod download

# Container의 Working Directory로 Code Copy
COPY . .

# Build
RUN go build -o main .

FROM nvidia/cuda:11.8.0-base-ubuntu22.04

WORKDIR /server
COPY --from=builder /server/main .

# Execute
CMD ["./main"]