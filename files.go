package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func CheckForFile(fileName string) error {
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		_, err := os.Create(fileName)
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateSlice(fileName string) []EmailRequestBody {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	var data []EmailRequestBody

	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil
	}

	return data
}
