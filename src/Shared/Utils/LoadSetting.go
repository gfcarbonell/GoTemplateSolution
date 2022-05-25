package Utils

import (
	"TemplateSolution/src/Shared/Models"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"
	"log"
	"os"
)

const (
	Environment                    = "ENVIRONMENT"
	Extension                      = "yaml"
	DevelopmentEnvironment         = "DEVELOPMENT"
	QAEnvironment                  = "QA"
	ProductionEnvironment          = "PRODUCTION"
	FilenameDevelopmentEnvironment = "setting.dev"
	FilenameQAEnvironment          = "setting.qa"
	FilenameProductionEnvironment  = "setting.prod"
)

func Init(path string, filename string, extension string) {
	viper.SetConfigName(filename)

	switch extension {
	case "env":
		viper.SetConfigType("env")
	case "yml":
		viper.SetConfigType("yaml")
	case "json":
		viper.SetConfigType("json")
	default:
		viper.SetConfigType("yaml")
	}

	viper.AddConfigPath(path)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}
}

func GetValueSetting(path string, filename string, extension, key string) interface{} {
	Init(path, filename, extension)
	return viper.Get(key)
}

var SETTING Models.Setting

func LoadSetting(path string) {
	Setting := &SETTING

	var filename string
	value := GetValueSetting(path, "app", Extension, Environment)

	switch value {
	case DevelopmentEnvironment:
		filename = FilenameDevelopmentEnvironment
	case QAEnvironment:
		filename = FilenameQAEnvironment
	case ProductionEnvironment:
		filename = FilenameProductionEnvironment
	default:
		filename = FilenameDevelopmentEnvironment
	}

	log.Println("Init Environment:", value)

	Init(path, filename, Extension)

	if err := viper.Unmarshal(&Setting); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	spew.Dump(SETTING)
}
