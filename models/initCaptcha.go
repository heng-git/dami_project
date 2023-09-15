package models

import (
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"

	"go-micro.dev/v4/client"
	"go-micro.dev/v4/registry"
)

var CaptchaClient client.Client

func init() {

	//配置consul
	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)
	// Create service
	srv := micro.NewService(
		micro.Registry(consulReg),
	)
	srv.Init()
	CaptchaClient = srv.Client()

}
