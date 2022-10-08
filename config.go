package main

import (
	"github.com/spf13/viper"
)

type EndpointConfig struct {
	Status int    `mapstructure:"status"`
	Body   string `mapstructure:"body"`
	Delay  int    `mapstructure:"delay"`
}

type ResourceConfig struct {
	Target   string `mapstructure:"target"`
	Step     string `mapstructure:"step"`
	Interval int    `mapstructure:"interval"`
}

type ServerConfig struct {
	Port          int `mapstructure:"port"`
	StartDelay    int `mapstructure:"startDelay"`
	ShutdownDelay int `mapstructure:"shutdownDelay"`
}

type Config struct {
	Server    ServerConfig   `mapstructure:"server"`
	Health    EndpointConfig `mapstructure:"health"`
	Readiness EndpointConfig `mapstructure:"readiness"`
	Resources struct {
		Cpu     ResourceConfig `mapstructure:"cpu"`
		Memory  ResourceConfig `mapstructure:"memory"`
		Storage ResourceConfig `mapstructure:"storage"`
	} `mapstructure:"resources"`
}

func LoadConfig() (config Config, err error) {
	vp := viper.New()
	vp.AddConfigPath(".")
	vp.SetConfigName("config") // Register config file name (no extension)
	vp.SetConfigType("yaml")   // Look for specific type
	err = vp.ReadInConfig()
	if err != nil {
		return
	}

	err = vp.Unmarshal(&config)
	if err != nil {
		return
	}
	return
}
