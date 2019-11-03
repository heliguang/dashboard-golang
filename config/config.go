package config

import (
	"io/ioutil"
	"runtime"

	"gopkg.in/yaml.v2"

	"dashboard/logger"
)

var Conf = &Config{}

type Route struct {
	Path      string  `yaml:"path" json:"path,omitempty"`
	Icon      string  `yaml:"icon" json:"icon,omitempty"`
	Title     string  `yaml:"title" json:"title,omitempty"`
	Component string  `yaml:"component" json:"component,omitempty"`
	Redirect  string  `yaml:"redirect" json:"redirect,omitempty"`
	Children  []Route `yaml:"children" json:"children,omitempty"`
}

type Config struct {
	ApiAddress string  `yaml:"api_address"`
	WebAddress string  `yaml:"web_address"`
	Routes     []Route `yaml:"routes"`
}

func init() {

	// TODO: 这里不够优雅
	// 给yaml路径默认是服务端部署时的的路径
	// 当系统环境是mac或者windows时，
	// 认为是开发环境，将yaml文件路径指定为绝对路径，
	// 是为了防止在package中的TestCase执行时，
	// 找不到yaml文件的报错
	yamlPath := "config.yaml"
	switch runtime.GOOS {
	case "darwin":
		yamlPath = ""
	case "windows":
		yamlPath = "D:/project/golang/src/dashboard/config.yaml"
	}
	logger.Info("load yaml from:" + yamlPath)
	yamlFile, err := ioutil.ReadFile(yamlPath)

	if err != nil {
		logger.Error("load config file error[" + err.Error() + "]")
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, Conf)
	if err != nil {
		logger.Error("unmarshal config file error[" + err.Error() + "]")
		panic(err)
	}
	logger.Info("load config success")
}
