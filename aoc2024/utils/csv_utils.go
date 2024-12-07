package utils

import (
	"encoding/csv"
	"os"
)

func ReadCsvFile(path string, readAll bool) ([][]string, error) {
	pwd, _ := os.Getwd()
	f, err := os.Open(pwd + path)
	res := [][]string{}
	if err != nil {
		//handle error after log configuration
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	if readAll {

		data, err := csvReader.ReadAll()
		if err != nil {
			//handle error after log configuration
		}
		return data, nil
	} else {

		for {
			data, err := csvReader.Read()
			if err != nil {
				if err.Error() == "EOF" {
					break
				}
				//handle error after log configuration
			}
			res = append(res, data)
		}
	}
	return res, nil
}
