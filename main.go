package main

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
	"github.com/tealeg/xlsx"
	"time"
)

type Error struct {
	Status int `json:"status"`
	Code int `json:"code"`
	Message string `json:"message"`
}

func main() {
	var x map[string]Error
	data, _ := ioutil.ReadFile("./errors.json")

	err := json.Unmarshal(data, &x)
	if err != nil {
		panic(err)
	}
	// Convert the map to Excel file
	WriteErrorsInExcel(x)
}

func WriteErrorsInExcel(errors map[string]Error){
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Errors")
	// Creating the new file
	if err != nil {
		panic(err)
	}
	for index, errorMsg := range errors {
		row := sheet.AddRow()
		errorNameCell := row.AddCell()
		codeCell := row.AddCell()
		statusCell := row.AddCell()
		messageCell := row.AddCell()
		errorNameCell.SetValue(index)
		codeCell.SetValue(errorMsg.Code)
		statusCell.SetValue(errorMsg.Status)
		messageCell.SetValue(errorMsg.Message)
	}
	// Saving the file with time to prevent colision
	err = file.Save(fmt.Sprintf("errors-%v.xlsx", time.Now().Unix()))
	if err != nil {
		panic(err)
	}
}