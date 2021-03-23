package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"4d63.com/tz"
	"github.com/spf13/viper"
)

// Key is used as a config key
type Key string

// These constants hold all config value keys
const (
	ServerEnv      Key = `server.env`
	ServerTimezone Key = `server.timezone`

	ServerHttpIP   Key = `server.http.ip`
	ServerHttpPort Key = `server.http.port`
	ServerHttpCORS Key = `server.http.cors`

	LogsEnabled Key = `logs.enabled`
	LogsLevel   Key = `logs.log_level`

	LogsFileEnabled Key = `logs.file.enabled`
	LogsFileName    Key = `logs.file.name`
	LogsFilePath    Key = `logs.file.path`

	DatabaseName     Key = `database.name`
	DatabaseHost     Key = `database.host`
	DatabaseAuth     Key = `database.auth`
	DatabaseUsername Key = `database.username`
	DatabasePassword Key = `database.password`

	RedisEnabled  Key = `redis.enabled`
	RedisHost     Key = `redis.host`
	RedisPort     Key = `redis.port`
	RedisAuth     Key = `redis.auth`
	RedisPassword Key = `redis.password`

	JwtAccessSecret  Key = `jwt.ACCESS_SECRET`
	JwtRefreshSecret Key = `jwt.REFRESH_SECRET`
)

func InitConfig() {
	viper.SetConfigName("config") // config file name without extension
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config/") // config file path
	viper.AutomaticEnv()             // read value ENV variable

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

}

// GetString returns a string config value
func (k Key) GetString() string {
	return viper.GetString(string(k))
}

// GetBool returns a bool config value
func (k Key) GetBool() bool {
	return viper.GetBool(string(k))
}

// GetInt returns an int config value
func (k Key) GetInt() int {
	return viper.GetInt(string(k))
}

// GetInt64 returns an int64 config value
func (k Key) GetInt64() int64 {
	return viper.GetInt64(string(k))
}

var timezone *time.Location

func GetTimeZone() *time.Location {
	if timezone == nil {
		loc, err := tz.LoadLocation(ServerTimezone.GetString())
		if err != nil {
			fmt.Printf("Error parsing time zone: %s", err)
			os.Exit(1)
		}
		timezone = loc
	}
	return timezone
}

// Set sets a value
func (k Key) Set(i interface{}) {
	viper.Set(string(k), i)
}
