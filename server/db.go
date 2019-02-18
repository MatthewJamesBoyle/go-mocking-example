package server

import (
	"database/sql"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type SQLstorer interface {
	SaveResult(result string) error
}

type MongoStorer interface {
	saveSum(numberOne int64, numberTwo int64) error
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

func (n NoSQL) saveSum(numberOne int64, numberTwo int64) error {
	return nil
}

func (SQL) SaveResult(result string) error {
	return nil
}
