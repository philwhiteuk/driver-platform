package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type driver struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func parseDriverData() ([]driver, error) {
	dat, err := ioutil.ReadFile("./drivers.csv")
	if err != nil {
		return nil, err
	}
	r := csv.NewReader(strings.NewReader(string(dat)))

	records, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("Invalid CSV %s", err)
	}

	var drivers []driver
	for _, r := range records[1:] {
		i, err := strconv.Atoi(r[0])
		if err != nil {
			return nil, fmt.Errorf("Invalid Driver ID %s", err)
		}
		drivers = append(drivers, driver{ID: i, Name: r[1]})
	}

	return drivers, nil
}
