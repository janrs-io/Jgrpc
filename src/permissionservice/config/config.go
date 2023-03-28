package config

import (
	"net/http"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

// Config Service config
type Config struct {
	Database Database `json:"database" yaml:"database"`
	Grpc     Grpc     `json:"grpc" yaml:"grpc"`
	Http     Http     `json:"http" yaml:"http"`
	Client   Client   `json:"client" yaml:"client"`
	Env      string   `json:"env" yaml:"env"`
}

// NewConfig Initial service's config
func NewConfig(cfg string) *Config {

	if cfg == "" {
		panic("load config file failed.config file can not be empty.")
	}

	viper.SetConfigFile(cfg)

	// Read config file
	if err := viper.ReadInConfig(); err != nil {
		panic("read config failed.[ERROR]=>" + err.Error())
	}
	conf := &Config{}
	// Assign the overloaded configuration to the global
	if err := viper.Unmarshal(conf); err != nil {
		panic("assign config failed.[ERROR]=>" + err.Error())
	}

	return conf

}

// Database Database config
type Database struct {
	Driver              string `json:"driver" yaml:"driver"`
	Host                string `json:"host" yaml:"host"`
	Port                int    `json:"port" yaml:"port"`
	UserName            string `json:"username" yaml:"username"`
	Password            string `json:"password" yaml:"password"`
	Database            string `json:"database" yaml:"database"`
	Charset             string `json:"charset" yaml:"charset"`
	MaxIdleCons         int    `json:"maxIdleCons" yaml:"maxIdleCons"`
	MaxOpenCons         int    `json:"maxOpenCons" yaml:"maxOpenCons"`
	LogMode             string `json:"logMode" yaml:"logMode"`
	EnableFileLogWriter bool   `json:"enableFileLogWriter" yaml:"enableFileLogWriter"`
	LogFilename         string `json:"logFilename" yaml:"logFilename"`
}

// Grpc Grpc server config
type Grpc struct {
	Host   string `json:"host" yaml:"host"`
	Port   string `json:"port" yaml:"port"`
	Name   string `json:"name" yaml:"name"`
	Server *grpc.Server
}

// Http Http server config
type Http struct {
	Host   string `json:"host" yaml:"host"`
	Port   string `json:"port" yaml:"port"`
	Name   string `json:"name" yaml:"name"`
	Server *http.Server
}

func initConfig(env string, conf *Config) {
	switch env {

	}

}

func initDevConfig() {

}

func initTestConfig() {

}

func initProdConfig() {

}
