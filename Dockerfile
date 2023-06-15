FROM golang:1.18-alpine AS build

WORKDIR /app

COPY . .

RUN apk add --no-cache gcc musl-dev

RUN go fmt ./...
RUN go vet ./...
RUN go mod tidy
RUN go build -gcflags "all=-N -l" -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/main .

CMD ["./main"]
