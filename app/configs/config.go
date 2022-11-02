package configs

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/kelseyhightower/envconfig"
)

const accessTokenTemplate = "%s:%s@tcp(%s)/%s"

type databaseConfig struct {
	User string `envconfig:"DB_USER" default:"funcy"`
	Pass string `envconfig:"DB_PASSWORD" default:"userpass"`
	IP   string `envconfig:"DB_IP" default:"mysql"`
	Name string `envconfig:"DB_NAME" default:"funcy"`
}

func GetServerPort() string {
	var addr string
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}
	// 接続情報は以下のように指定する.
	// user:password@tcp(host:port)/database
	flag.StringVar(&addr, "addr", ":"+port, "tcp host:port to connect")
	flag.Parse()
	return addr
}

func GetDBConnectionInfo() string {
	/* ===== データベースへ接続する. ===== */
	var config databaseConfig
	if err := envconfig.Process("", &config); err != nil {
		log.Fatal("Unable to connect to DB(Insufficient variables)")
	}
	log.Println("Starting db")
	return fmt.Sprintf(accessTokenTemplate, config.User, config.Pass, config.IP, config.Name)
}
