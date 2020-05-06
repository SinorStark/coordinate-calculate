package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// LocateData all location info 城市位置信息结构体
type LocateData struct {
	Countary  string  `json:"country"`
	Province  string  `json:"province"`
	City      string  `json:"city"`
	Area      string  `json:"area"`
	Longitude float64 `json:"lng"`
	Latitude  float64 `json:"lat"`
}

// type LocateDatas struct {
// 	Info []LocateData
// }

// NewLocateData ...
func NewLocateData(countary string, province string, city string,
	area string, longitude float64, latitude float64) *LocateData {
	return &LocateData{countary, province, city, area, longitude, latitude}
}

// GetAllDataFromFile ...
func GetAllDataFromFile() ([]LocateData, error) {
	// datas := make([]LocateData, 100, 300)
	var err error
	var data []LocateData
	fmt.Println(err)
	if dtFile, fileErr := ioutil.ReadFile("data.json"); fileErr != nil {
		fmt.Println(fileErr.Error())
		// return nil, fileErr
	} else {
		if json.Valid(dtFile) {
			if err := json.Unmarshal(dtFile, &data); err != nil {
				fmt.Println(err.Error())
			}
		}
		// if len(datas) >= cap(datas) {
		// 	datas = append(datas, data)
		// }
	}
	return data, nil
}
