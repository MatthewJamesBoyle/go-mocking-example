package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type CalcHandler struct {
	Summer   SumStorer
	Resulter ResultStorer
	logger   SuperComplexLogger
}

type SumRequest struct {
	First  string `json:"first"`
	Second string `json:"second"`
}

func NewHandler(Resulter ResultStorer, Summer SumStorer, logger SuperComplexLogger) (*CalcHandler, error) {

	return &CalcHandler{
		Summer:   Summer,
		Resulter: Resulter,
		logger:   logger,
	}, nil

}

func (c *CalcHandler) AddNumbers(w http.ResponseWriter, req *http.Request) {
	var r SumRequest
	decoder := json.NewDecoder(req.Body)

	err := decoder.Decode(&r)
	if r.First == "" || r.Second == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	first, err := strconv.ParseInt(r.First, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	second, err := strconv.ParseInt(r.Second, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result := first + second
	//using testify, we get a nil pointer exception here if we don't set expectations.
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
