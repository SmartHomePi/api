package initialize

import (
	"github.com/SmartHomePi/api/pkg/config"
	"github.com/SmartHomePi/api/pkg/logger"
)

func LightInit() {
	// Init the config
	config.InitConfig()
	logger.InitLogger()

	// redid.InitRedis()
	// keyvalue.InitStorage()
	// log.InitLogger()
}

// InitEngines intializes all db connections
func InitEngines() {

}

func FullInit() {
	LightInit()
	InitEngines()
	logger.Debug("Initialization done..")
}
