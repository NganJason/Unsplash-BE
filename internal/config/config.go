package config

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/NganJason/Unsplash-BE/pkg/cerr"
	"github.com/NganJason/Unsplash-BE/pkg/clog"
)

var (
	GlobalConfig *Config
)

type Config struct {
	UnsplashDB *Database `json:"unsplash_db"`
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

const configFilePath = "../internal/config/config.json"

func initConfigs() error {
	fetchConfigFromFile(configFilePath, &GlobalConfig)

	if GlobalConfig.UnsplashDB == nil {
		clog.Fatal(context.Background(), "fail to init unsplashDB")
	}

	initDBs()

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
