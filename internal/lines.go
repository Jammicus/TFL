package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/kyokomi/emoji"
)

var apiURL = "https://api.tfl.gov.uk/line/mode/tube/status"

// API Structs

type Lines struct {
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	ModeName     string       `json:"modeName"`
	Distruptions string       `json:"distruptions"`
	Created      time.Time    `json:"created"`
	Modified     time.Time    `json:"modified"`
	LineStatuses []lineStatus `json:"lineStatuses"`
}

type lineStatus struct {
	ID                   int               `json:"id"`
	LineID               string            `json:"lineId"`
	StatusSeverity       int               `json:"statusSeverity"`
	StatusSevDescription string            `json:"statusSeverityDescription"`
	Reason               string            `json:"reason"`
	Created              createdTime       `json:"created"`
	Period               []validityPeriods `json:"validityPeriods"`
}

type validityPeriods struct {
	From  fromTime `json:"fromDate"`
	To    toTime   `json:"toDate"`
	IsNow bool     `json:"isNow"`
}

type disrutpion struct {
	Category            string `json:"category"`
	CategoryDescription string `json:"categoryDescription"`
	Description         string `json:"description"`
	ClosureText         string `json:"closureText"`
}

//Post Processing structs
type LineSummary struct {
	Name   string
	Status string
	Info   string
}

type fromTime struct {
	time.Time
}

type toTime struct {
	time.Time
}

type createdTime struct {
	time.Time
}

func convertTime(layout string, value []byte) time.Time {
	strInput := strings.Trim(string(value), `"`)

	newTime, err := time.Parse(layout, strInput)
	if err != nil {
		log.Fatal("Error occured" + err.Error())
	}
	return newTime
}

// Handle weird time format in validityPeriods
func (t *fromTime) UnmarshalJSON(input []byte) error {
	t.Time = convertTime("2006-01-02T15:04:05Z", input)
	return nil
}

// Handle weird time format in validityPeriods
func (t *toTime) UnmarshalJSON(input []byte) error {
	t.Time = convertTime("2006-01-02T15:04:05Z", input)
	return nil
}

// Handle weird time format in lineStatus
func (t *createdTime) UnmarshalJSON(input []byte) error {
	t.Time = convertTime("2006-01-02T15:04:05", input)
	return nil
}

//GetLines Gets Details Of All London Underground Lines
func getLines() []Lines {
	var lines []Lines

	resp, err := http.Get(apiURL)
	if err != nil {
		log.Fatal("Error occured" + err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error occured" + err.Error())
		os.Exit(1)
	}

	err = json.Unmarshal(responseData, &lines)

	if err != nil {
		log.Fatal("Error occured " + err.Error())
		os.Exit(1)
	}
	return lines
}

func determineStatus(lines []Lines) []LineSummary {
	var lineSum = make([]LineSummary, len(lines))
	for index, item := range lines {
		// var status string
		status, err := DecodeStatus(item.LineStatuses[0].StatusSeverity, item.ModeName)
		if err != nil {
			log.Fatal("Error occured" + err.Error())
			os.Exit(1)
		}
		lineSum[index] = LineSummary{item.Name, status, item.LineStatuses[0].Reason}
	}

	return lineSum
}

func PrintStatus() {
	for _, line := range determineStatus(getLines()) {
		emoji.Println(":metro:", " ", line.Name, " ", line.Status, " ", line.Info)
	}
}
