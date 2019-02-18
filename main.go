package main

import (
	"fmt"
	"github.com/MatthewJamesBoyle/mocking-example/server"
	"net/http"
	"os"
)

func main() {

	calcHandler, err := server.NewHandler(
		"host=localhost user=summer password=password1 dbname=summer sslmode=disable",
		"mongodb://root:root@localhost",
	)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	http.ListenAndServe(":8085", server.Routes(calcHandler))
}
