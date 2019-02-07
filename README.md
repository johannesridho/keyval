# Keyval

An app to set and get key value using Go and Redis

### How to Run

1. Install Go
2. Install Redis
3. Run `go get github.com/johannesridho/keyval`
3. Update .env based on your need
4. Run commands: 
    - `vgo build`
    - `go run github.com/johannesridho/keyval`
5. Server will run at localhost:8080 by default

### Endpoints

#### Set Keyval
```
Request:

POST /keyvals
Body (Content-Type application/json):
{
	"key": "key1",
	"value": "val1"
}
```

```
Response:
{
    "key": "key1",
    "value": "val1"
}
```

#### Get Keyval
```
Request:

Get /keyvals/key1
```

```
Response:
{
    "key": "key1",
    "value": "val1"
}
```
