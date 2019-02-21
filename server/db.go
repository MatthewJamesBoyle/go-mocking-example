package server

import (
	"database/sql"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type ResultStorer interface {
	Save(result string) error
}

type SumStorer interface {
	Save(numberOne int64, numberTwo int64) error
}

type NoSQL struct {
	mongo *mongo.Client
}

type SQL struct {
	db *sql.DB
}

func NewSQL(db *sql.DB) SQL {
	return SQL{
		db: db,
	}
}
func NewNoSQL(db *mongo.Client) NoSQL {
	return NoSQL{
		mongo: db,
	}
}

func (n NoSQL) Save(numberOne int64, numberTwo int64) error {
	return nil
}

func (s SQL) Save(result string) error {
	return nil
}
