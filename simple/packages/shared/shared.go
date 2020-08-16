package shared

import (
	"github.com/spf13/viper"
)

type Configs struct {
	Hostname string `json:"hostname" mapstructure:"hostname"`
}

func InitConfigs(path string, env string) (*Configs, error) {
	config := &Configs {}

	v := viper.New()

	v.SetConfigType("env")

	envPath := path + "/.env"
	v.SetConfigFile(envPath)

	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	if env == "" {
		env = v.GetString("environment")
	}

	if env == "" {
		env = "development"
	}

	err = v.Unmarshal(config)
	if err != nil {
		return nil, err
	}

	envPath = path + "/.env." + env
	v.SetConfigFile(envPath)
	err = v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = v.Unmarshal(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}