package api

import (
	"encoding/json"
	"fmt"
	"github.com/kyokomi/emoji"
	"io/ioutil"
	"net/http"
	"os"
)

// TODO: API is returning weird date format. Need to figure out how to handle with time.Time.
// 	fmt.Println(lines[0].Created.Format(time.ANSIC))

var apiURL = "https://api.tfl.gov.uk/line/mode/tube/status"

// API Structs

type Lines struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	ModeName     string `json:"modeName"`
	Distruptions string `json:"distruptions"`
	// Created      time.Time    `json:"created"`
	// Modified     time.Time    `json:"modified"`
	LineStatuses []lineStatus `json:"lineStatuses"`
}

type lineStatus struct {
	ID                   int    `json:"id"`
	LineID               string `json:"lineId"`
	StatusSeverity       int    `json:"statusSeverity"`
	StatusSevDescription string `json:"statusSeverityDescription"`
	Reason               string `json:"reason"`
	// Created              time.Time       `json:"created"`
	Period []validityPeriods `json:"validityPeriods"`
}

type validityPeriods struct {
	// From  time.Time `json:"fromDate"`
	// To    time.Time `json:"toDate"`
	IsNow bool `json:"isNow"`
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

func determineStatus(lines []Lines) []LineSummary {
	var lineSum = make([]LineSummary, len(lines))
	for index, item := range lines {
		// var status string
		status, err := DecodeStatus(item.LineStatuses[0].StatusSeverity, item.ModeName)
		if err != nil {
			fmt.Print("Error occured" + err.Error())
			os.Exit(1)
		}
		lineSum[index] = LineSummary{item.Name, status, item.LineStatuses[0].Reason}
	}

	return lineSum
}

func PrintStatus() {
	for _, line := range determineStatus(GetLines()) {
		emoji.Println(":metro:", " ", line.Name, " ", line.Status, " ", line.Info)
	}
}
