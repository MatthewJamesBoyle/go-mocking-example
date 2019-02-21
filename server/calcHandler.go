package server

import (
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"net/http"
	"strconv"
)

type CalcHandler struct {
	Summer   SumStorer
	Resulter ResultStorer
	logger   SuperComplexLogger
}

func NewHandler(Resulter ResultStorer, Summer SumStorer, logger SuperComplexLogger) (*CalcHandler, error) {

	return &CalcHandler{
		Summer:   Summer,
		Resulter: Resulter,
		logger:   logger,
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

	err = c.Summer.Save(first, second)
	if err != nil {
		c.logger.SuperLog("SUPER LOG")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = c.Resulter.Save(string(result))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *CalcHandler) MinusNumbers(w http.ResponseWriter, req *http.Request) {

}
