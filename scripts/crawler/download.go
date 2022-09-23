package main

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func downloadFile(URL, fileName string) (string, error) {
	response, err := http.Get(URL)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return "", errors.New("Received non 200 response code")
	}

	file, err := os.Create(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return "", err
	}

	absPath, err := filepath.Abs(fileName)
	if err != nil {
		return "", err
	}

	return absPath, nil
}

type Data struct {
	Photo          interface{}
	Path           string
	ProfileImgPath string
}

func writeJsonData(output []Data) {
	jsonFile, _ := os.Open("data.json")
	jsonBytes, _ := ioutil.ReadAll(jsonFile)

	var initialData []Data
	json.Unmarshal(jsonBytes, &initialData)

	finalOutput := append(initialData, output...)

	file, _ := json.MarshalIndent(finalOutput, "", " ")
	_ = ioutil.WriteFile("data.json", file, 0644)
}
