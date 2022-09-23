package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/hbagdi/go-unsplash/unsplash"
)

type SeedDataRequest struct {
	EmailAddress *string `json:"email_address"`
	Password     *string `json:"password"`
	Username     *string `json:"username"`
	FirstName    *string `json:"first_name"`
	LastName     *string `json:"last_name"`
	ImageDesc    *string `json:"image_desc"`
}

const (
	remoteURL = "http://localhost:8082/api/data/seed"
)

func main() {
	data := getJsonData()

	for idx, info := range data {
		photo := info.Photo

		req := SeedDataRequest{
			EmailAddress: StrPtr(*photo.Photographer.Username + "@gmail.com"),
			Password:     StrPtr("123"),
			Username:     photo.Photographer.Username,
			FirstName:    photo.Photographer.FirstName,
			LastName:     photo.Photographer.LastName,
			ImageDesc:    photo.Description,
		}

		bytes, _ := json.Marshal(req)

		values := map[string]io.Reader{
			"img":         mustOpen(info.Path),
			"data":        strings.NewReader(string(bytes)),
			"profile_img": mustOpen(info.ProfileImgPath),
		}

		err := Upload(remoteURL, values)
		if err != nil {
			log.Println(err.Error())
			continue
		}

		log.Println(fmt.Sprintf("Processed img %d", idx))
	}
}

type Data struct {
	Photo          unsplash.Photo `json:"Photo"`
	Path           string         `json:"Path"`
	ProfileImgPath string         `json:"ProfileImgPath"`
}

// func fetchPhotoInfos() {
// 	data := getJsonData()
// 	photos := make([]unsplash.Photo, 0)

// 	for _, info := range data {
// 		infoByte, _ := json.Marshal(info.Photo)

// 	}
// 	photoInfo := data[0].Photo

// 	infoByte, _ := json.Marshal(photoInfo)
// 	var info unsplash.Photo

// 	json.Unmarshal(infoByte, &info)
// 	fmt.Println(*info.Photographer.Name)
// }
func getJsonData() []Data {
	jsonFile, _ := os.Open("../crawler/data.json")
	jsonBytes, _ := ioutil.ReadAll(jsonFile)

	var data []Data
	json.Unmarshal(jsonBytes, &data)

	return data
}
