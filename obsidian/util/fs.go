package util

import (
	"log"
	"os"
)

func GetPwd() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		return "/tmp"
	}
	return dir
}
