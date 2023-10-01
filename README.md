# Project content-service

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

