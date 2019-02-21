package main

import (
	"database/sql"
	"fmt"
	"github.com/MatthewJamesBoyle/mocking-example/server"
	"github.com/mongodb/mongo-go-driver/mongo"
	"net/http"
	"os"
)

func main() {

	sqlConString := ""
	noSqlDbString := ""

	db, _ := sql.Open("postgres", sqlConString)

	mongo, _ := mongo.NewClient(noSqlDbString)

	resulter := server.NewNoSQL(mongo)
	summer := server.NewSQL(db)

	calcHandler, err := server.NewHandler(summer, resulter, server.MattsLogger{})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	http.ListenAndServe(":8085", server.Routes(calcHandler))
}
