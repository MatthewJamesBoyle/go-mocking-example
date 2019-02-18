package main

import (
	"github.com/MatthewJamesBoyle/mocking-example/server"
	"net/http"
)

func main() {
	h := server.CalcHandler{}
	http.ListenAndServe(":8085", server.Routes(&h))
}
