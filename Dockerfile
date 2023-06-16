FROM golang:1.18-alpine AS build

WORKDIR /app

COPY . .

RUN apk add --no-cache gcc musl-dev

RUN go fmt ./...
RUN go vet ./...
RUN go mod tidy

# Build the Go project
RUN go build -o main .

CMD ["./main"]
