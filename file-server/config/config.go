package config

import (
	"flag"
	"os"
)

func GetServerPort() string {
	var addr string
	port := os.Getenv("PORT")
	if port == "" {
		port = "3004"
	}
	// 接続情報は以下のように指定する.
	// user:password@tcp(host:port)/database
	flag.StringVar(&addr, "addr", ":"+port, "tcp host:port to connect")
	flag.Parse()
	return addr
}
