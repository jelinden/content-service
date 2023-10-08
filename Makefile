BINARY_NAME=content-service
LOCALBINARYNAME=${BINARY_NAME}-arm

build:
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux main.go
	GOARCH=arm64 GOOS=darwin go build -o ${BINARY_NAME}-arm main.go

run-local:
	make clean
	cd frontend && npm run build && cd ..
	go build -o ${LOCALBINARYNAME} main.go
	./${LOCALBINARYNAME}

clean:
	go clean
	rm ${BINARY_NAME}-linux || true
	rm ${BINARY_NAME}-arm || true

help: 
	@echo "\n\t use 'make build' to make a build \n\t 'make run-local' to run on arm mac \n\t 'make clean' to clean\n"