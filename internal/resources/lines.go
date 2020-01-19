package lines

import (
	"fmt"
	"net/http"
	"os"
	"time"
	"encoding/json"
	"io/ioutil"
)
// Handling Date.
// 	fmt.Println(lines[0].Created.Format(time.ANSIC))


var apiURL = "https://api.tfl.gov.uk/line/mode/tube/status"

type Lines struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	ModeName     string    `json:"modeName"`
	Distruptions string    `json:"distruptions"`
	Created      time.Time `json:"created"`
	modified     time.Time `json:"modified"`
	lineStatuses lineStatus `json:"lineStatuses"`
}

type lineStatus struct {
	ID                   int             `json:"id"`
	LineID               string          `json:"lineId"`
	StatusSeverity       int             `json:"statusSeverity"`
	StatusSevDescription string          `json:"statusSeverityDescription"`
	Reason               string          `json:"statusSeverity"`
	Created              time.Time       `json:"created"`
	Period               validityPeriods `json:"validityPeriods"`
}

type validityPeriods struct {
	from  time.Time `json:"fromDate"`
	to    time.Time `json:"toDate"`
	isNow bool      `json:"isNow"`
}

type disrutpion struct {
	Category            string `json:"category"`
	CategoryDescription string `json:"categoryDescription"`
	Description         string `json:"description"`
	ClosureText         string `json:"closureText"`
}
//GetLines Gets Details Of All London Underground Lines
func GetLines() []Lines {
	var lines []Lines

	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Print("Error occured" + err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print("Error occured" + err.Error())
		os.Exit(1)
	}
	err = json.Unmarshal(responseData, &lines)
	
	if err != nil {
		fmt.Print("Error occured " + err.Error())
		os.Exit(1)
	}
	
	return lines
}
