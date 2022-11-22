package service

import (
	"code.sajari.com/docconv"
	"fmt"
	"log"
)

func ConvertToHtml(path string) (string, error) {
	res, err := docconv.ConvertPath(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)

	return "", err
}
