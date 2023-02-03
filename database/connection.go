package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/mauFade/go-postgres/configs"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDatabase()

	stringConnection := fmt.Sprintf(
		"host=%sport=%suser=%sdatabase=%spassword:%ssslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Database, conf.Password,
	)

	connection, err := sql.Open("postgres", stringConnection)

	if err != nil {
		panic(err)
	}

	err = connection.Ping()

	return connection, err
}
