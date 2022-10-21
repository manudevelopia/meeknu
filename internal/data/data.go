package data

import (
	"github.com/jackc/pgx/v5/pgxpool"
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
