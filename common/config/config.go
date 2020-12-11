package config

import (
	"bytes"
	"github.com/huahuayu/address-generator/common/dir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

type Configuration struct {
	Env    string `yaml:"env"`
	Server struct {
		Port        string `yaml:"port"`
		TimezoneLoc string `yaml:"timezoneLoc"`
		GinMode     string `yaml:"ginMode"`
	} `yaml:"server"`
	Db struct {
		User       string `yaml:"user"`
		Pass       string `yaml:"pass"`
		Host       string `yaml:"host"`
		Port       string `yaml:"port"`
		Name       string `yaml:"name"`
		MaxConnect int    `yaml:"maxConnect"`
		MaxIdle    int    `yaml:"maxIdle"`
		ShowSQL    bool   `yaml:"showSql"`
	} `yaml:"db"`
	Redis struct {
		Host string `yaml:"host"`
		Pass string `yaml:"pass"`
		Db   int    `yaml:"db"`
	} `yaml:"redis"`
	Log struct {
		Path  string    `yaml:"path"`
		Level log.Level `yaml:"level"`
	} `yaml:"log"`
}

var App *Configuration

func Init(configPath string, env string) {
	configType := "yml"
	v := viper.New()
	v.SetConfigType(configType)

	// if config path passed by flag, use it
	if configPath != "" {
		getConfigByFlag(configPath, configType, v)
		return
	}

	// otherwise read config from project dir
	configDir := dir.Root + string(os.PathSeparator) + "config"
	log.Info(configDir)
	getDefaultConfig(v, configDir)

	// read env specific configs. priority: env flag > config/app.yml setting
	getEnvConfig(env, v, configDir)

	unmarshalConfig(v)
}

func getConfigByFlag(configPath string, configType string, v *viper.Viper) {
	confDir := path.Dir(configPath)
	name := path.Base(configPath)
	suffix := path.Ext(configPath)

	if suffix != "."+configType {
		log.Fatal("wrong config file type")
	}
	v.SetConfigName(strings.TrimSuffix(name, suffix))
	v.AddConfigPath(confDir)
	err := v.ReadInConfig()
	if err != nil {
		log.Fatal("read config file error: ", err)
	}
	log.Info("active config: ", configPath)
	unmarshalConfig(v)
}

func getDefaultConfig(v *viper.Viper, configDir string) {
	v.SetConfigName("app")
	v.AddConfigPath(configDir)
	err := v.ReadInConfig()
	if err != nil {
		log.Fatal("read default config file error: ", err)
	}

	// load all settings into default setting
	for k, val := range v.AllSettings() {
		v.SetDefault(k, val)
	}
}

func getEnvConfig(envFlag string, v *viper.Viper, configDir string) {
	env := getEnv(envFlag, v)
	file, err := ioutil.ReadFile(configDir + string(os.PathSeparator) + "app-" + env + ".yml")
	if err != nil {
		log.Fatal("read env config file error", err)
	}
	err = v.ReadConfig(bytes.NewReader(file))
	if err != nil {
		log.Fatal("read config file error: ", err)
	}
}

func getEnv(envFlag string, v *viper.Viper) string {
	envInConfigFile := v.GetString("env")
	env := envInConfigFile
	if envFlag != "" {
		env = strings.ToLower(envFlag)
	}
	v.Set("env", env)
	log.Info("active config: ", env)
	return env
}

func unmarshalConfig(v *viper.Viper) {
	err := v.Unmarshal(&App)
	if err != nil {
		log.Fatal("wrong config file format")
	}
}
