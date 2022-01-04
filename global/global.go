package global

import (
	"github.com/qingchunyibeifangzongle/initialization_go_project/config"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	GVA_CONFIG config.Server
	GVA_VP     *viper.Viper

	BlackCache local_cache.Cache

	GVA_DB *gorm.DB
	MYSQL  *gorm.DB
)
