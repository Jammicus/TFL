package lines

import (
	"http"
	"time"
)

var apiURL = "https://api.tfl.gov.uk/line/mode/tube/status"


type Lines struct {
	ID string `json:"id"`
	Name Array `json:"name"`
	ModeName string `json:"modeName"`
	Distruptions [] `json:"distruptions"`
	Created time.Time `json:"created"`
	modified time.Time `json:"modified"`
	// Check, should be an array of structs
	lineStatuses lineStatus	`json:"lineStatuses"`
	routeSections [] `json:"serviceTypes"`
}

type lineStatus struct {
	ID	int `json:"id"`
	LineID string `json:"lineId"`
	StatusSeverity int `json:"statusSeverity"`
	StatusSevDescription string `json:"statusSeverityDescription"`
	Reason string `json:"statusSeverity"`
	Created time.Time `json:"created"`
	Period ValidityPeriods `json:"validityPeriods"`
}

type ValidityPeriods struct {
	from time.Time `json:"fromDate"`
	to time.Time `json:"toDate"`
	isNow bool `json:"isNow"`
}

type disrutpion struct {
	Category string `json:"category"`
	CategoryDescription string `json:"categoryDescription"`
	Description string `json:"description"`
	AffectedRoutes [] `json:"affectedRoutes"`
	AffectedStops []  `json:"affectedStops"`
	ClosureText string `json:"closureText"`
}



func getLineStatus(string apiURL, string line) {
	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Print("Error occured" + err.Error())
		os.exit(1)
	}

	data,err := ioutil.ReadAll(resp.body)
	if err != nil {
		fmt.Print("Error occured" + err.Error())
		os.exit(1)
	}
	fmt.Println(data)

}
