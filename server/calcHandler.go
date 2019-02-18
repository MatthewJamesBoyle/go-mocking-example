package server

import (
	"database/sql"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type CalcHandler struct {
	SqlDB   SQLstorer
	NoSqlDB MongoStorer
}

func NewHandler(sqlConString string, noSqlDbString string) (*CalcHandler, error) {
	if sqlConString == "" || noSqlDbString == "" {
		return nil, errors.New("no con string")
	}
	db, err := sql.Open("postgres", sqlConString)
	if err != nil {
		return nil, err
	}

	mongo, err := mongo.NewClient(noSqlDbString)
	if err != nil {
		return nil, err
	}

	return &CalcHandler{
		SqlDB:   NewSQL(db),
		NoSqlDB: NewNoSQL(mongo),
	}, nil

}

func (c *CalcHandler) AddNumbers(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	fVar := vars["first"]
	sVar := vars["second"]
	if fVar == "" || sVar == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	first, err := strconv.ParseInt(fVar, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	second, err := strconv.ParseInt(sVar, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result := first + second

	err = c.NoSqlDB.saveSum(first, second)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = c.SqlDB.SaveResult(string(result))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *CalcHandler) MinusNumbers(w http.ResponseWriter, req *http.Request) {

}
