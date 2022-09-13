package main

import (
	"context"
	"fmt"
	"log"
)

const (
	imgPath     = "./images/%s.png"
	jsonPath    = "./data.json"
	numRequests = 20
	startPage   = 1
)

func main() {
	ctx := context.Background()
	dm := NewPhotoDM(ctx)

	pageNum := startPage
	globalData := make([]Data, 0)

	for i := 0; i < numRequests; i++ {
		photos, resp, err := dm.GetPhotos(pageNum, 100)
		if err != nil {
			log.Println(err)
			log.Printf("current page num: %d", pageNum)
			return
		}

		pageNum = resp.NextPage

		for idx, photo := range *photos {
			fmt.Println(
				fmt.Sprintf("Page num: %d | Img idx: %d", pageNum-1, idx),
			)

			fileName := fmt.Sprintf(imgPath, *photo.ID)
			absPath, err := downloadFile(photo.Links.Download.String(), fileName)
			if err != nil {
				log.Println(err)
				log.Printf("current page num: %d", pageNum)
				continue
			}

			d := Data{
				Photo: photo,
				Path:  absPath,
			}

			globalData = append(globalData, d)
		}

		writeJsonData(globalData)
		globalData = []Data{}
	}

	defer writeJsonData(globalData)
}
