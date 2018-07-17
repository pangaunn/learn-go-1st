package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type StarwarPeople struct {
	Name      string
	Height    string
	Weight    int    `json:"mass,string"`
	HairColor string `json:"hair_color"`
	Films     []string
}

func main() {
	get()
	post()
}

func post() {
	params := strings.NewReader(`{"hello":"aaaa"}`)
	resp, _ := http.Post("http://192.168.0.244:5001/", "application/json", params)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func get() {
	url := "https://swapi.co/api/people/1"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)

	var starwarPeople StarwarPeople
	json.Unmarshal(body, &starwarPeople)
	fmt.Println(starwarPeople)

	blob, _ := json.Marshal(starwarPeople)
	fmt.Println(string(blob))
}
