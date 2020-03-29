package controllers

// Operations about object
type HealthCheckController struct {
	BaseController
}

// @Title getEpBySrv
// @Description get dolphin health status
// @Success 200 {string}  active
// @router /health [get]
func (c *HealthCheckController) Get() {
	data := map[string]interface{}{
		"active":  true,
		"details": "",
	}
	c.RespJson(CtlResp{}.SetData(data))

	return
}
