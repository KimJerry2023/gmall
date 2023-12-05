package main

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"gmall/conf"
	"gmall/loading"
	"gmall/routes"
)

func main() {
	fmt.Println("OSS Go SDK version: ", oss.Version)
	conf.Init()
	loading.Loading()
	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)
}
