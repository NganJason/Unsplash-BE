package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

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

func main() {
	// var remoteURL string
	// remoteURL = "http://localhost:8082/api/data/seed"

	// data := SeedDataRequest{
	// 	EmailAddress: StrPtr("test1"),
	// 	Password: StrPtr("123"),
	// 	Username: StrPtr("test1"),
	// 	FirstName: StrPtr("testFirst"),
	// 	LastName: StrPtr("testLast"),
	// 	ImageDesc: StrPtr("test_image"),
	// }

	// bytes, _ := json.Marshal(data)

	// values := map[string]io.Reader{
	//     "img":  mustOpen("_9b4PK2na7c.png"),
	//     "data": strings.NewReader(string(bytes)),
	// }

	// err := Upload(remoteURL, values)
	// if err != nil {
	//     panic(err)
	// }
	data := getJsonData()
	photoInfo := data[0].Photo

	infoByte, _ := json.Marshal(photoInfo)
	var info unsplash.Photo

	json.Unmarshal(infoByte, &info)
	fmt.Println(*info.Photographer.Name)

}

type Data struct {
	Photo interface{}
	Path  string
}

func getJsonData() []Data {
	jsonFile, _ := os.Open("../crawler/data.json")
	jsonBytes, _ := ioutil.ReadAll(jsonFile)

	var data []Data
	json.Unmarshal(jsonBytes, &data)

	return data
}
