package api

// https://api.tfl.gov.uk/Line/Meta/Severity
//TODO Make inputs lowercase?? OR at least consistent

import (
	"errors"
	"fmt"
)

var cableCar = map[int]string{
	0:  "Special Service",
	1:  "Closed",
	2:  "No Service",
	3:  "No Service",
	4:  "Planned Closure",
	5:  "Part Closure",
	6:  "Severe Delays",
	7:  "Reduced Service",
	8:  "Bus Service",
	9:  "Minor Delays",
	10: "Good Service",
	11: "Part Closed",
	12: "Exit only",
	13: "No Step Free Access",
	14: "Change of Frequency",
	15: "Diverted",
	16: "Not Running",
	17: "Issues Reported",
	18: "No Issues",
	19: "Information",
}

var cycle = map[int]string{
	0:  "Special Service",
	1:  "Closed",
	2:  "No Service",
	3:  "No Service",
	4:  "Planned Closure",
	5:  "Part Closure",
	6:  "Severe Delays",
	7:  "Reduced Service",
	8:  "Bus Service",
	9:  "Minor Delays",
	10: "Good Service",
	11: "Part Closed",
	12: "Exit only",
	13: "No Step Free Access",
	14: "Change of Frequency",
	15: "Diverted",
	16: "Not Running",
	17: "Issues Reported",
	18: "No Issues",
	19: "Information",
}

var dlr = map[int]string{
	0:  "Special Service",
	1:  "Closed",
	2:  "No Service",
	3:  "No Service",
	4:  "Planned Closure",
	5:  "Part Closure",
	6:  "Severe Delays",
	7:  "Reduced Service",
	8:  "Bus Service",
	9:  "Minor Delays",
	10: "Good Service",
	11: "Part Closed",
	12: "Exit only",
	13: "No Step Free Access",
	14: "Change of Frequency",
	15: "Diverted",
	16: "Not Running",
	17: "Issues Reported",
	18: "No Issues",
	19: "Information",
	20: "Service Closed",
}

var overground = map[int]string{
	0:  "Special Service",
	1:  "Closed",
	2:  "No Service",
	3:  "No Service",
	4:  "Planned Closure",
	5:  "Part Closure",
	6:  "Severe Delays",
	7:  "Reduced Service",
	8:  "Bus Service",
	9:  "Minor Delays",
	10: "Good Service",
	11: "Part Closed",
	12: "Exit only",
	13: "No Step Free Access",
	14: "Change of Frequency",
	15: "Diverted",
	16: "Not Running",
	17: "Issues Reported",
	18: "No Issues",
	19: "Information",
	20: "Service Closed",
}

var tflrail = map[int]string{
	0:  "Special Service",
	1:  "Closed",
	2:  "No Service",
	3:  "No Service",
	4:  "Planned Closure",
	5:  "Part Closure",
	6:  "Severe Delays",
	7:  "Reduced Service",
	8:  "Bus Service",
	9:  "Minor Delays",
	10: "Good Service",
	11: "Part Closed",
	12: "Exit only",
	13: "No Step Free Access",
	14: "Change of Frequency",
	15: "Diverted",
	16: "Not Running",
	17: "Issues Reported",
	18: "No Issues",
	19: "Information",
	20: "Service Closed",
}

var bus = map[int]string{
	0:  "Special Service",
	1:  "Closed",
	2:  "No Service",
	3:  "No Service",
	4:  "Planned Closure",
	5:  "Part Closure",
	6:  "Severe Delays",
	7:  "Reduced Service",
	8:  "Bus Service",
	9:  "Minor Delays",
	10: "Good Service",
	11: "Part Closed",
	12: "Exit only",
	13: "No Step Free Access",
	14: "Change of Frequency",
	15: "Diverted",
	16: "Not Running",
	17: "Issues Reported",
	18: "No Issues",
	19: "Information",
	20: "Service Closed",
}

//Boats?
var riverBus = map[int]string{
	0:  "Special Service",
	1:  "Closed",
	2:  "No Service",
	3:  "No Service",
	4:  "Planned Closure",
	5:  "Part Closure",
	6:  "Severe Delays",
	7:  "Reduced Service",
	8:  "Bus Service",
	9:  "Minor Delays",
	10: "Good Service",
	11: "Part Closed",
	12: "Exit only",
	13: "No Step Free Access",
	14: "Change of Frequency",
	15: "Diverted",
	16: "Not Running",
	17: "Issues Reported",
	18: "No Issues",
	19: "Information",
	20: "Service Closed",
}

var riverTour = map[int]string{
	0:  "Special Service",
	1:  "Closed",
	2:  "No Service",
	3:  "No Service",
	4:  "Planned Closure",
	5:  "Part Closure",
	6:  "Severe Delays",
	7:  "Reduced Service",
	8:  "Bus Service",
	9:  "Minor Delays",
	10: "Good Service",
	11: "Part Closed",
	12: "Exit only",
	13: "No Step Free Access",
	14: "Change of Frequency",
	15: "Diverted",
	16: "Not Running",
	17: "Issues Reported",
	18: "No Issues",
	19: "Information",
	20: "Service Closed",
}

var tram = map[int]string{
	0:  "Special Service",
	1:  "Closed",
	2:  "No Service",
	3:  "No Service",
	4:  "Planned Closure",
	5:  "Part Closure",
	6:  "Severe Delays",
	7:  "Reduced Service",
	8:  "Bus Service",
	9:  "Minor Delays",
	10: "Good Service",
	11: "Part Closed",
	12: "Exit only",
	13: "No Step Free Access",
	14: "Change of Frequency",
	15: "Diverted",
	16: "Not Running",
	17: "Issues Reported",
	18: "No Issues",
	19: "Information",
	20: "Service Closed",
}

var tube = map[int]string{
	0:  "Special Service",
	1:  "Closed",
	2:  "No Service",
	3:  "No Service",
	4:  "Planned Closure",
	5:  "Part Closure",
	6:  "Severe Delays",
	7:  "Reduced Service",
	8:  "Bus Service",
	9:  "Minor Delays",
	10: "Good Service",
	11: "Part Closed",
	12: "Exit only",
	13: "No Step Free Access",
	14: "Change of Frequency",
	15: "Diverted",
	16: "Not Running",
	17: "Issues Reported",
	18: "No Issues",
	19: "Information",
	20: "Service Closed",
}

// DecodeStatus Takes a status code and mode of Transport.
// Returns the meaning behind the status code
func DecodeStatus(code int, modeName string) (string, error) {
	switch modeName {
	case "cableCar":
		value, exists := cableCar[code]
		if exists == true {
			return value, nil
		}
		return "", errors.New("Invalid Error Code " + fmt.Sprint(code))

	case "cycle":
		value, exists := cycle[code]
		if exists == true {
			return value, nil
		}
		return "", errors.New("Invalid Error Code " + fmt.Sprint(code))

	case "dlr":
		value, exists := dlr[code]
		if exists == true {
			return value, nil
		}
		return "", errors.New("Invalid Error Code " + fmt.Sprint(code))

	case "overground":
		value, exists := overground[code]
		if exists == true {
			return value, nil
		}
		return "", errors.New("Invalid Error Code " + fmt.Sprint(code))

	case "tflRail":
		value, exists := tflrail[code]
		if exists == true {
			return value, nil
		}
		return "", errors.New("Invalid Error Code " + fmt.Sprint(code))

	case "bus":
		value, exists := bus[code]
		if exists == true {
			return value, nil
		}
		return "", errors.New("Invalid Error Code " + fmt.Sprint(code))

	case "riverBus":
		value, exists := riverBus[code]
		if exists == true {
			return value, nil
		}
		return "", errors.New("Invalid Error Code " + fmt.Sprint(code))

	case "riverTour":
		value, exists := riverTour[code]
		if exists == true {
			return value, nil
		}
		return "", errors.New("Invalid Error Code " + fmt.Sprint(code))

	case "tram":
		value, exists := tram[code]
		if exists == true {
			return value, nil
		}
		return "", errors.New("Invalid Error Code " + fmt.Sprint(code))

	case "tube":
		value, exists := tube[code]
		if exists == true {
			return value, nil
		}
		return "", errors.New("Invalid Error Code " + fmt.Sprint(code))

	default:
		return "", errors.New("Invalid Type " + modeName)
	}
}
