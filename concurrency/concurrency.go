package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ResponseApi struct {
	Help    string `json:"help"`
	Success bool   `json:"success"`
	Result  struct {
		ResourceID string `json:"resource_id"`
		Fields     []struct {
			Type string `json:"type"`
			ID   string `json:"id"`
		} `json:"fields"`
		Records []struct {
			ID            int    `json:"_id"`
			Sex           string `json:"sex"`
			NoOfGraduates string `json:"no_of_graduates"`
			TypeOfCourse  string `json:"type_of_course"`
			Year          string `json:"year"`
		} `json:"records"`
		Links struct {
			Start string `json:"start"`
			Next  string `json:"next"`
		} `json:"_links"`
		Limit int `json:"limit"`
		Total int `json:"total"`
	} `json:"result"`
}

func main() {
	url := "https://data.gov.sg/api/action/datastore_search?resource_id=eb8b932c-503c-41e7-b513-114cffbe2338&limit=1"
	response := getData(url)
	fmt.Println(response)

}

func getData(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Can't read data body")
	}

	json.Unmarshal(data, ResponseApi{})
	return string(data)
}

// func worker() {

// }

// func job() {

// }
