package db

import (
	"database/sql"
	"fmt"

	"github.com/aryelzx/api-go/configs"
	_ "github.com/lib/pq" //driver connection postgres
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()

	//connection string
	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s slmode=disabled",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Database,
	)

	conn, err := sql.Open("postgres", sc)
	if err != nil {
		panic(err)
	}
	err = conn.Ping()

	return conn, err
}
