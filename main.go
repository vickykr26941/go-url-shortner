package main

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/vickykumar/url_shortner/cmd/server"
)

func main() {
	jwt.MarshalSingleStringAsArray = false
	server.Execute()
}
