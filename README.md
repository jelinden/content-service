# Project content-service

The purpose of this project is to make a service which has a per user possibility to 
* register
* login
* save an API key for themselfs
* save key/value pairs and
* get key/value pairs with the key through an API with an API key

## run backend

`go build && ./content-service`

or

`go run main.go`

## environment variable JWT_KEY

For JWT we need a secret key with witch we make the token for the user/browser.

Setting the environment variable can be done for example with:

`export JWT_KEY=your-secret-key-here`

Then this environment variable is used in authorize.go

`secretKey = os.Getenv("JWT_KEY")`

