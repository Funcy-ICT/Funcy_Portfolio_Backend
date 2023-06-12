package main

import (
	"backend/file-server/config"
	"backend/file-server/pkg"
)

func main() {

	//サーバを起動
	addr := config.GetServerPort()
	pkg.Server.Run(addr)
}
