package api

import (
	"github.com/tigren/go-tour/init_trap/config"
)

var (
	timeoutInMs int64
	apiEnabled  = true
)

//IsEnabled ...
func IsEnabled() bool {
	return apiEnabled
}

func init() {
	config.Init()
	if config.Config.ApiConf == nil {
		apiEnabled = false
		return
	}

	timeoutInMs = config.Config.ApiConf.TimeoutInMs
}
