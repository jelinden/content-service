BINARY_NAME=content-service

build:
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux main.go
	GOARCH=arm64 GOOS=darwin go build -o ${BINARY_NAME}-arm main.go

run-local:
	go build -o ${BINARY_NAME} main.go
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}-linux
	rm ${BINARY_NAME}-arm

help: 
	@echo "\n\t use 'make build' to make a build \n\t 'make run-local' to run on arm mac \n\t 'make clean' to clean\n"