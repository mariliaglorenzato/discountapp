package config

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

const (
	ServerEnv  = "SERVER_ENVIRONMENT"
	DevEnv     = "development"
	StagingEnv = "staging"
	TestEnv    = "test"
	ProdEnv    = "production"

	eventsDiscounts = "events_discounts"
)

func GetBasePath() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Dir(b)
}

func SetConfigs(env string) {
	if env == "" {
		env = DevEnv
	}

	viper.SetDefault(ServerEnv, env)
}

func GetEnv() string {
	return viper.GetString(ServerEnv)
}

func LoadEventConfigFromFile() []byte {
	eventConfigFileName := GetBasePath() + "/" + eventsDiscounts + "." + GetEnv() + ".json"

	jsonEvents, err := os.Open(eventConfigFileName)
	if err != nil {
		panic("Error at opening EventsConfig: " + err.Error())
	}

	byteEvents, err := ioutil.ReadAll(jsonEvents)
	if err != nil {
		panic("Error at reading EventsConfig: " + err.Error())
	}
	return byteEvents
}
