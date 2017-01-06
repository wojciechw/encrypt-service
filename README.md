# Encrypt Server

## Dependencies

* go get github.com/gorilla/mux
* go get github.com/golang/protobuf/proto
* go get golang.org/x/crypto/nacl/secretbox

## Usage

* go get github.com/wojciechw/encrypt-service
* go build github.com/wojciechw/encrypt-service/server
* ./server

Server runs on port :12345
Port can be changed by option -port=45677

* go build github.com/wojciechw/encrypt-service/client
* client -e -id=123 -data="Test text"
* client -d -id=123 -key=key

Client connects to default server address: http://localhost:12345
A server address can be changed by option -addr=http://encrypt-server.com:45677

For easy use key is returned in hexadecimal encoding

## Running tests

* go test -v github.com/wojciechw/encrypt-service/server
* go test -v github.com/wojciechw/encrypt-service/storage

