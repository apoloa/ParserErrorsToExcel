package main

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
)

type Error struct {
	Status int `json:"status"`
	Code int `json:"code"`
	Message string `json:"message"`
}

func main() {
	var x map[string]Error
	data, _ := ioutil.ReadFile("./errors.json")

	decoder := json.Unmarshal(data, &x)
	fmt.Println(decoder)
	fmt.Println(x)

}