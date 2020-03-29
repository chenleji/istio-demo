package main

import (
	"github.com/astaxie/beego"
	_ "github.com/chenleji/istio-demo/routers"
	"github.com/chenleji/nautilus/helper"
	"os"
)

const (
	AppHealthCheckURL = "/health"
	KeyAppName        = "app"
)

func main() {
	// consul
	registerService()

	beego.Run()
}

func registerService() {
	if beego.BConfig.RunMode == beego.DEV {
		appName := os.Getenv(KeyAppName)
		if len(appName) == 0 {
			panic("invalid app name")
		}

		err := helper.Consul{}.New().RegistryService(
			appName,
			helper.Utils{}.GetAppPort(),
			AppHealthCheckURL)
		if err != nil {
			panic(err)
		}

	}
}
