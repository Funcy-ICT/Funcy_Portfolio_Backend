package main

import (
	"backend/configs"
	"backend/pkg"
	"backend/pkg/model/dao"
	"log"
)

func main() {
	//DBのコネクションを確率
	err := dao.Init()
	if err != nil {
		log.Fatal(err)
	}

	//サーバを起動
	addr := configs.GetServerPort()
	pkg.Server.Run(addr)
}
