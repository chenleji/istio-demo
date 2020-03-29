package controllers

import (
	"fmt"
	"github.com/chenleji/nautilus/helper"
	"net/http"
	"os"
	"strings"
)

const (
	KeyMode         = "mode"
	KeyAppName      = "app"
	KeyVersion      = "version"
	KeyNextServices = "next_services"

	ValueStubMode    = "stub"
	ValueSwitchMode  = "switch"
	ValueIngressMode = "ingress"
)

type MainController struct {
	BaseController
}

func (c *MainController) Get() {
	runMode := os.Getenv(KeyMode)
	nextServices := os.Getenv(KeyNextServices)
	serviceList := strings.Split(nextServices, ",")
	appName := os.Getenv(KeyAppName)
	if len(appName) == 0 {
		panic("invalid app name")
	}
	version := os.Getenv(KeyVersion)

	if runMode == ValueStubMode {
		data := map[string]interface{}{
			"service":     appName,
			"nextService": "-",
			"msg":         "version:" + version,
		}

		c.RespJson(CtlResp{}.
			SetCode(http.StatusOK).
			SetMsg("success").
			SetData(data))
		return
	} else {
		dataList := make([]interface{}, 0)

		for _, service := range serviceList {
			resp := new(CtlResp)
			header := map[string]string{
				"head": service,
			}

			err := helper.HttpClient{Service: service}.GetWithHeader(resp, "/", header, []string{})

			if err != nil { // fail
				dataItem := map[string]interface{}{
					"service":     appName,
					"nextService": fmt.Sprintf("%s", service),
					"msg":         fmt.Sprintf("exception: %s \n", err.Error()),
				}
				dataList = append(dataList, dataItem)
			} else { // success
				if resp.HttpCode == http.StatusOK {
					dataItem := map[string]interface{}{
						"service":     appName,
						"nextService": fmt.Sprintf("%s", service),
						"msg":         resp.Body,
					}
					dataList = append(dataList, dataItem)
				} else {
					dataItem := map[string]interface{}{
						"service":     appName,
						"nextService": fmt.Sprintf("%s", service),
						"msg":         fmt.Sprintf("response code: %d \n", resp.HttpCode),
					}
					dataList = append(dataList, dataItem)
				}
			}
		}

		if runMode == ValueIngressMode {
			c.Data["Website"] = "ljchen.net"
			c.Data["Email"] = "chenleji@gmail.com"
			c.Data["RespContent"] = dataList
			c.TplName = "get.tpl"
		} else {
			c.RespJson(CtlResp{}.
				SetCode(http.StatusOK).
				SetMsg("success").
				SetData(dataList))
		}
	}
}
