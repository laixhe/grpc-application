package conf

import (
	"fmt"

	"kratos-application/api/gen/config/capp"
	"kratos-application/api/gen/config/clog"
)

// Checking 检查
func Checking() {
	// app
	if cc.App.Version == "" {
		cc.App.Version = "v0.1"
	}
	if cc.App.Mode == "" {
		cc.App.Mode = capp.ModeType_debug.String()
	}
	fmt.Println("app", cc.App)
	// log
	if cc.Log.Run == "" {
		cc.Log.Run = clog.RunType_console.String()
	}
	if cc.Log.Level == "" {
		cc.Log.Level = clog.LevelType_debug.String()
	}
	if cc.Log.Run == clog.RunType_file.String() {
		if cc.Log.Path == "" {
			cc.Log.Path = "logs.log"
		}
	}
}
