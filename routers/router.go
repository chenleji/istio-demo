package routers

import (
	"github.com/astaxie/beego"
	"github.com/chenleji/istio-demo/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/health", &controllers.HealthCheckController{})
}
