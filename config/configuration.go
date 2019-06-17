package config

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/viper"
)

// config base info.
const (
	ConfigType = "yml"
	ConfigPath = "config"
	EnvDev     = "development"
)

type UserServiceConfig struct {
	Host string
	Port int
}

type EventServiceConfig struct {
	Host string
	Port int
}

// Configurations all configurations of app
type Configurations struct {
	UserService  *UserServiceConfig
	EventService *EventServiceConfig
}

func (c *Configurations) unmarshalConfig(files []*configFile) error {
	var err error
	for _, f := range files {
		v := viper.New()
		v.SetConfigName(f.filename)
		v.AddConfigPath(f.path)
		v.AutomaticEnv()
		if err = v.ReadInConfig(); err != nil {
			return err
		}
		if err = v.Unmarshal(&c); err != nil {
			return err
		}
	}
	return nil
}

type configFile struct {
	path     string
	filename string
}

// NewConfigurations get configurations instance
func NewConfigurations(env string) (*Configurations, error) {
	viper.SetConfigType(ConfigType)
	configFiles := []*configFile{
		&configFile{path: getConfigPath(), filename: env},
	}

	c := &Configurations{}
	err := c.unmarshalConfig(configFiles)

	return c, err
}

func getConfigPath() string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	return path.Join(pwd, ConfigPath)
}
