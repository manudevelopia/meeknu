package data

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"sync"
)

var (
	data *Data
	once sync.Once
)

type Data struct {
	DB *pgxpool.Pool
}

func initDB() {
	db := getDataBaseConnection()
	data = &Data{
		DB: db,
	}
}

func New() *Data {
	once.Do(initDB)
	return data
}

func (d Data) Query(query string) pgx.Rows {
	rows, err := d.DB.Query(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	return rows
}
