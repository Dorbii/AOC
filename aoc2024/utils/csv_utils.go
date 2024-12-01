package utils

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ReadCsvFile(path string) ([][]string, error) {
	pwd, _ := os.Getwd()
	f, err := os.Open(pwd + path)
	fmt.Println(pwd + path)
	if err != nil {
		//handle error after log configuration
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		//handle error after log configuration
	}
	return data, nil
}
