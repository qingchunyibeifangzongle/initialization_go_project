package main

import (
	"github.com/qingchunyibeifangzongle/initialization_go_project/core"
	"github.com/qingchunyibeifangzongle/initialization_go_project/global"
	"github.com/qingchunyibeifangzongle/initialization_go_project/initialize"
)

func main() {
	global.GVA_VP = core.Viper()

	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	//fmt.Println(global.GVA_DB)
}
