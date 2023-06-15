build:
	go fmt ./...
	go vet ./...
	go mod tidy
	go build -gcflags "all=-N -l" -o bin/go-template ./

run:
	go fmt ./...
	go vet ./...
	go mod tidy
	go run ./main.go

docker-build:
	docker build -t go-template .

docker-run:
	docker-compose up

clean:
	rm -rf bin/
