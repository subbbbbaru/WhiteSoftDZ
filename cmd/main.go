package main

import (
	"log"
	"os"

	"github.com/subbbbbaru/whitesoftdz/menu"
	"github.com/subbbbbaru/whitesoftdz/record"
)

func readData(url string) ([]byte, error) {
	expectedData, err := os.ReadFile(url)
	if err != nil {
		return nil, err
	}
	return expectedData, nil
}

func main() {
	args := os.Args
	if len(args) != 2 {
		log.Fatal("Wrong arguments!")
	}

	data, err := readData(args[1])
	if err != nil {
		log.Fatal(err.Error())
	}

	recordSlice := record.RecordSlice{}
	err = recordSlice.FromJson(data)
	if err != nil {
		log.Fatal(err.Error())
	}
	recordMap := recordSlice.ToMap()

	menu.MainMenu(&recordMap)
}
