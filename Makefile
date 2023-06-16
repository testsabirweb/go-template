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

docker-build-debug:
	docker build . --tag go-template-debug --file Dockerfile.debug

docker-run:
	docker-compose up mysql go-template

docker-run-debug:
	docker-compose up -d mysql
	docker run --publish 80:80 --publish 3000:3000 --name debug-server go-template-debug

clean:
	rm -rf bin/
	
#docker run --publish 80:80 --publish 3000:3000 --name debug-server go-template-debug
#docker rm $(docker ps -a -q)