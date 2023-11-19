# Project content-service

The purpose of this project is to make a service which has a per user possibility to 
* register
* login
* save a API key for themselfs
* save key/value pairs and
* get key/value pairs with the key through an API with an API key


## Usage

Register to https://content-service.jelinden.fi

Login

Make a space https://content-service.jelinden.fi/space

and add content after which you can get content for example to your web site from

https://content-service.jelinden.fi/api/space/:spaceID/entries?token=YOUR_TOKEN_FROM_PROFILE

Remember cors-issues though, it is better to proxy the content through a server on your domain.

## Run backend and frontend with makefile

`make run-local`

does the following:

```
run-local:
	make clean
 	cd frontend && npm run build && cd .. // run build makes a production build and copies it to be served from go server
 	go build -o content-service-arm main.go // build the binary
 	./content-service-arm  // run the binary
 ```

## Run backend tests

`go test ./...`


## Run backend independently

`go build && ./content-service`

or

`go run main.go`

## Run frontend independently

`cd frontend && npm install && npm run start`


## Environment variable JWT_KEY when running the server

For JWT we need a secret key with witch we make the token for the user/browser.

Setting the environment variable can be done for example with:

`export JWT_KEY=your-secret-key-here`

Then this environment variable is used in authorize.go

`secretKey = os.Getenv("JWT_KEY")`

## Tech used

Backend: Golang, router github.com/julienschmidt/httprouter, database sqlite, JWT tokens (github.com/golang-jwt/jwt/v5)

Frontend: Typescript, React, React router

## Database

Database is embedded and creates a file of its own to the directory the app is running. As long as nothing is done to it, the content is preserved.
Restarting the service doesn't do anything to the file.
