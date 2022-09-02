package config

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/NganJason/BE-template/pkg/cerr"
)

var (
	GlobalConfig *Config
)

type Config struct {
	SampleDB *Database `json:"sample_db"`
}

func GetConfig() *Config {
	return GlobalConfig
}

func InitGlobalConfig() error {
	GlobalConfig = new(Config)

	err := initConfigs()
	if err != nil {
		return err
	}

	return nil
}

func initConfigs() error {
	return nil
}

func fetchConfigFromFile(
	filePath string,
	configStruct interface{},
) error {
	configFile, _ := os.Open(filePath)
	defer configFile.Close()

	decoder := json.NewDecoder(configFile)
	err := decoder.Decode(configStruct)
	if err != nil {
		return cerr.New(
			fmt.Sprintf("load configs err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}

	return nil
}
