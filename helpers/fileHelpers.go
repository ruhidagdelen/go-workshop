package helpers

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ReadFile() (error, records [][]string) {
	file, err := os.Open("/Users/ruhidagdelen/Projects/personal/" +
		"GO/go-workshop/static_data/top500domains.csv")
	if err != nil {
		fmt.Println("we could not able to read file")
		return nil, nil
	}
	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()

	return nil, data
}
