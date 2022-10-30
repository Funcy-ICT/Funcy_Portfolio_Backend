package configs

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"

	"github.com/cenkalti/backoff"

	// blank import for MySQL driver
	_ "github.com/go-sql-driver/mysql"
)

// Driver名
const DriverName = "mysql"

// Conn 各repositoryで利用するDB接続(Connection)情報
var Conn *sqlx.DB
var DBConnectionInfo string

func Init() (*sqlx.DB, error) {
	DBConnectionInfo = GetDBConnectionInfo()
	conn, err := createDBConnection()
	if err != nil {
		if err = dbConnectionBackoff(); err != nil {
			return nil, err
		}
	}
	log.Println("Successfull DB Connection")
	return conn, err
}

func createDBConnection() (*sqlx.DB, error) {
	var err error
	Conn, err = sqlx.Open(DriverName, DBConnectionInfo)
	if err != nil {
		return nil, err
	}
	if err = Conn.Ping(); err != nil {
		return nil, err
	}
	return Conn, nil
}

func dbConnectionBackoff() error {
	b := backoff.WithMaxRetries(backoff.NewExponentialBackOff(), 7)
	err := backoff.Retry(Conn.Ping, b)
	if err != nil {
		log.Println(fmt.Errorf("Faild create connection"))
		return err
	}
	return err
}
