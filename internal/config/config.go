package config

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

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

const configFilePath = "./internal/config/config.json"

func initConfigs() error {
	fetchConfig()

	if GlobalConfig.UnsplashDB == nil {
		clog.Fatal(context.Background(), "fail to init unsplashDB")
	}

	initDBs()

	return nil
}

func fetchConfig() {
	DBUrl := os.Getenv("DATABASE_URL")
	if DBUrl != "" {
		split := strings.Split(DBUrl, ":")
		dbUsername := split[0]
		dbPassword := split[1]
		dbHost := split[2]
		dbName := split[3]

		GlobalConfig.UnsplashDB = &Database{
			Username: dbUsername,
			Password: dbPassword,
			DBName:   dbName,
			Host:     dbHost,
		}
	} else {
		fetchConfigFromFile(configFilePath, &GlobalConfig)
	}

	if GlobalConfig == nil {
		log.Fatal("failed to fetch configs")
	}
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
