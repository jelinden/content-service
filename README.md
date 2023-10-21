# Project content-service

The purpose of this project is to make a service which has a per user possibility to 
* register
* login
* save an API key for themselfs
* save key/value pairs and
* get key/value pairs with the key through an API with an API key

## run backend and frontend with makefile

`make run-local`

does the following:

`run-local:
	make clean
	cd frontend && npm run build && cd .. // run build makes a production build and copies it to be served from go server
	go build -o content-service-arm main.go // build the binary
	./content-service-arm` // run the binary


## run backend

`go build && ./content-service`

or

`go run main.go`

## run frontend

`cd frontend && npm install && npm run start`


## environment variable JWT_KEY

For JWT we need a secret key with witch we make the token for the user/browser.

Setting the environment variable can be done for example with:

`export JWT_KEY=your-secret-key-here`

Then this environment variable is used in authorize.go

`secretKey = os.Getenv("JWT_KEY")`

