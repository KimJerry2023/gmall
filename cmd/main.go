package main

import (
	"gmall/conf"
	"gmall/loading"
	"gmall/routes"
)

func main() {
	conf.Init()
	loading.Loading()
	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)
}
