package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/MatthewJamesBoyle/mocking-example/server"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {

	sqlConString := ""
	noSqlDbString := ""

	db, _ := sql.Open("postgres", sqlConString)

	mongoClient, _ := mongo.Connect(options.Client().ApplyURI(noSqlDbString))
	defer mongoClient.Disconnect(context.Background())

	resulter := server.NewNoSQL(mongoClient)
	summer := server.NewSQL(db)

	calcHandler, err := server.NewHandler(summer, resulter, server.MattsLogger{})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	http.ListenAndServe(":8085", server.Routes(calcHandler))
}
