package api

import (
	"testing"
)

func TestCableCarValid(t *testing.T) {
	var testCases = []struct {
		input    int
		expected string
	}{
		{0, "Special Service"},
		{1, "Closed"},
		{2, "No Service"},
		{3, "No Service"},
		{4, "Planned Closure"},
		{5, "Part Closure"},
		{6, "Severe Delays"},
		{7, "Reduced Service"},
		{8, "Bus Service"},
		{9, "Minor Delays"},
		{10, "Good Service"},
		{11, "Part Closed"},
		{12, "Exit only"},
		{13, "No Step Free Access"},
		{14, "Change of Frequency"},
		{15, "Diverted"},
		{16, "Not Running"},
		{17, "Issues Reported"},
		{18, "No Issues"},
		{19, "Information"},
	}

	for _, test := range testCases {
		if item, _ := DecodeStatus(test.input, "cableCar"); item != test.expected {
			t.Error("Expected: ", test.expected, "But got: ", item)
		}
	}
}

func TestCableCarInvalid(t *testing.T) {
	var testCases = []struct {
		input int
	}{
		{-1},
		{20},
		{22},
		{23},
		{12021},
		{12023213},
	}

	for _, test := range testCases {
		if item, err := DecodeStatus(test.input, "cableCar"); err == nil {
			t.Error("Expected and error, but got: ", "But got: ", item)
		}
	}

}

func TestCycleValid(t *testing.T) {
	var testCases = []struct {
		input    int
		expected string
	}{
		{0, "Special Service"},
		{1, "Closed"},
		{2, "No Service"},
		{3, "No Service"},
		{4, "Planned Closure"},
		{5, "Part Closure"},
		{6, "Severe Delays"},
		{7, "Reduced Service"},
		{8, "Bus Service"},
		{9, "Minor Delays"},
		{10, "Good Service"},
		{11, "Part Closed"},
		{12, "Exit only"},
		{13, "No Step Free Access"},
		{14, "Change of Frequency"},
		{15, "Diverted"},
		{16, "Not Running"},
		{17, "Issues Reported"},
		{18, "No Issues"},
		{19, "Information"},
	}

	for _, test := range testCases {
		if item, _ := DecodeStatus(test.input, "cycle"); item != test.expected {
			t.Error("Expected: ", test.expected, "But got: ", item)
		}
	}
}

func TestCycleInvalid(t *testing.T) {
	var testCases = []struct {
		input int
	}{
		{-2},
		{20},
		{22},
		{23},
		{12021},
		{12023213},
	}

	for _, test := range testCases {
		if item, err := DecodeStatus(test.input, "cycle"); err == nil {
			t.Error("Expected and error, but got: ", "But got: ", item)
		}
	}
}

func TestDlrValid(t *testing.T) {
	var testCases = []struct {
		input    int
		expected string
	}{
		{0, "Special Service"},
		{1, "Closed"},
		{2, "No Service"},
		{3, "No Service"},
		{4, "Planned Closure"},
		{5, "Part Closure"},
		{6, "Severe Delays"},
		{7, "Reduced Service"},
		{8, "Bus Service"},
		{9, "Minor Delays"},
		{10, "Good Service"},
		{11, "Part Closed"},
		{12, "Exit only"},
		{13, "No Step Free Access"},
		{14, "Change of Frequency"},
		{15, "Diverted"},
		{16, "Not Running"},
		{17, "Issues Reported"},
		{18, "No Issues"},
		{19, "Information"},
		{20, "Service Closed"},
	}

	for _, test := range testCases {
		if item, _ := DecodeStatus(test.input, "dlr"); item != test.expected {
			t.Error("Expected: ", test.expected, "But got: ", item)
		}
	}
}

func TestDlrInvalid(t *testing.T) {
	var testCases = []struct {
		input int
	}{
		{-2},
		{22},
		{23},
		{12021},
		{12023213},
	}

	for _, test := range testCases {
		if item, err := DecodeStatus(test.input, "dlr"); err == nil {
			t.Error("Expected and error, but got: ", "But got: ", item)
		}
	}
}

func TestOvergroundValid(t *testing.T) {
	var testCases = []struct {
		input    int
		expected string
	}{
		{0, "Special Service"},
		{1, "Closed"},
		{2, "No Service"},
		{3, "No Service"},
		{4, "Planned Closure"},
		{5, "Part Closure"},
		{6, "Severe Delays"},
		{7, "Reduced Service"},
		{8, "Bus Service"},
		{9, "Minor Delays"},
		{10, "Good Service"},
		{11, "Part Closed"},
		{12, "Exit only"},
		{13, "No Step Free Access"},
		{14, "Change of Frequency"},
		{15, "Diverted"},
		{16, "Not Running"},
		{17, "Issues Reported"},
		{18, "No Issues"},
		{19, "Information"},
		{20, "Service Closed"},
	}

	for _, test := range testCases {
		if item, _ := DecodeStatus(test.input, "overground"); item != test.expected {
			t.Error("Expected: ", test.expected, "But got: ", item)
		}
	}
}

func TestOvergroundInvalid(t *testing.T) {
	var testCases = []struct {
		input int
	}{
		{-2},
		{22},
		{23},
		{12021},
		{12023213},
	}

	for _, test := range testCases {
		if item, err := DecodeStatus(test.input, "overground"); err == nil {
			t.Error("Expected and error, but got: ", "But got: ", item)
		}
	}
}

func TestTflRailValid(t *testing.T) {
	var testCases = []struct {
		input    int
		expected string
	}{
		{0, "Special Service"},
		{1, "Closed"},
		{2, "No Service"},
		{3, "No Service"},
		{4, "Planned Closure"},
		{5, "Part Closure"},
		{6, "Severe Delays"},
		{7, "Reduced Service"},
		{8, "Bus Service"},
		{9, "Minor Delays"},
		{10, "Good Service"},
		{11, "Part Closed"},
		{12, "Exit only"},
		{13, "No Step Free Access"},
		{14, "Change of Frequency"},
		{15, "Diverted"},
		{16, "Not Running"},
		{17, "Issues Reported"},
		{18, "No Issues"},
		{19, "Information"},
		{20, "Service Closed"},
	}

	for _, test := range testCases {
		if item, _ := DecodeStatus(test.input, "tflRail"); item != test.expected {
			t.Error("Expected: ", test.expected, "But got: ", item)
		}
	}
}

func TestTflRailInvalid(t *testing.T) {
	var testCases = []struct {
		input int
	}{
		{-2},
		{22},
		{23},
		{12021},
		{12023213},
	}

	for _, test := range testCases {
		if item, err := DecodeStatus(test.input, "tflrail"); err == nil {
			t.Error("Expected and error, but got: ", "But got: ", item)
		}
	}
}

func TestBusValid(t *testing.T) {
	var testCases = []struct {
		input    int
		expected string
	}{
		{0, "Special Service"},
		{1, "Closed"},
		{2, "No Service"},
		{3, "No Service"},
		{4, "Planned Closure"},
		{5, "Part Closure"},
		{6, "Severe Delays"},
		{7, "Reduced Service"},
		{8, "Bus Service"},
		{9, "Minor Delays"},
		{10, "Good Service"},
		{11, "Part Closed"},
		{12, "Exit only"},
		{13, "No Step Free Access"},
		{14, "Change of Frequency"},
		{15, "Diverted"},
		{16, "Not Running"},
		{17, "Issues Reported"},
		{18, "No Issues"},
		{19, "Information"},
		{20, "Service Closed"},
	}

	for _, test := range testCases {
		if item, _ := DecodeStatus(test.input, "bus"); item != test.expected {
			t.Error("Expected: ", test.expected, "But got: ", item)
		}
	}
}

func TestBusInvalid(t *testing.T) {
	var testCases = []struct {
		input int
	}{
		{-2},
		{22},
		{23},
		{12021},
		{12023213},
	}

	for _, test := range testCases {
		if item, err := DecodeStatus(test.input, "bus"); err == nil {
			t.Error("Expected and error, but got: ", "But got: ", item)
		}
	}
}

func TestRiverBusValid(t *testing.T) {
	var testCases = []struct {
		input    int
		expected string
	}{
		{0, "Special Service"},
		{1, "Closed"},
		{2, "No Service"},
		{3, "No Service"},
		{4, "Planned Closure"},
		{5, "Part Closure"},
		{6, "Severe Delays"},
		{7, "Reduced Service"},
		{8, "Bus Service"},
		{9, "Minor Delays"},
		{10, "Good Service"},
		{11, "Part Closed"},
		{12, "Exit only"},
		{13, "No Step Free Access"},
		{14, "Change of Frequency"},
		{15, "Diverted"},
		{16, "Not Running"},
		{17, "Issues Reported"},
		{18, "No Issues"},
		{19, "Information"},
		{20, "Service Closed"},
	}

	for _, test := range testCases {
		if item, _ := DecodeStatus(test.input, "riverBus"); item != test.expected {
			t.Error("Expected: ", test.expected, "But got: ", item)
		}
	}
}

func TestRiverBusInvalid(t *testing.T) {
	var testCases = []struct {
		input int
	}{
		{-2},
		{22},
		{23},
		{12021},
		{12023213},
	}

	for _, test := range testCases {
		if item, err := DecodeStatus(test.input, "riverBus"); err == nil {
			t.Error("Expected and error, but got: ", "But got: ", item)
		}
	}
}

func TestRiverTourValid(t *testing.T) {
	var testCases = []struct {
		input    int
		expected string
	}{
		{0, "Special Service"},
		{1, "Closed"},
		{2, "No Service"},
		{3, "No Service"},
		{4, "Planned Closure"},
		{5, "Part Closure"},
		{6, "Severe Delays"},
		{7, "Reduced Service"},
		{8, "Bus Service"},
		{9, "Minor Delays"},
		{10, "Good Service"},
		{11, "Part Closed"},
		{12, "Exit only"},
		{13, "No Step Free Access"},
		{14, "Change of Frequency"},
		{15, "Diverted"},
		{16, "Not Running"},
		{17, "Issues Reported"},
		{18, "No Issues"},
		{19, "Information"},
		{20, "Service Closed"},
	}

	for _, test := range testCases {
		if item, _ := DecodeStatus(test.input, "riverTour"); item != test.expected {
			t.Error("Expected: ", test.expected, "But got: ", item)
		}
	}
}

func TestRiverTourInvalid(t *testing.T) {
	var testCases = []struct {
		input int
	}{
		{-2},
		{22},
		{23},
		{12021},
		{12023213},
	}

	for _, test := range testCases {
		if item, err := DecodeStatus(test.input, "riverTour"); err == nil {
			t.Error("Expected and error, but got: ", "But got: ", item)
		}
	}
}

func TestTramValid(t *testing.T) {
	var testCases = []struct {
		input    int
		expected string
	}{
		{0, "Special Service"},
		{1, "Closed"},
		{2, "No Service"},
		{3, "No Service"},
		{4, "Planned Closure"},
		{5, "Part Closure"},
		{6, "Severe Delays"},
		{7, "Reduced Service"},
		{8, "Bus Service"},
		{9, "Minor Delays"},
		{10, "Good Service"},
		{11, "Part Closed"},
		{12, "Exit only"},
		{13, "No Step Free Access"},
		{14, "Change of Frequency"},
		{15, "Diverted"},
		{16, "Not Running"},
		{17, "Issues Reported"},
		{18, "No Issues"},
		{19, "Information"},
		{20, "Service Closed"},
	}

	for _, test := range testCases {
		if item, _ := DecodeStatus(test.input, "tram"); item != test.expected {
			t.Error("Expected: ", test.expected, "But got: ", item)
		}
	}
}

func TestTramInvalid(t *testing.T) {
	var testCases = []struct {
		input int
	}{
		{-2},
		{22},
		{23},
		{12021},
		{12023213},
	}

	for _, test := range testCases {
		if item, err := DecodeStatus(test.input, "tram"); err == nil {
			t.Error("Expected and error, but got: ", "But got: ", item)
		}
	}
}

func TestTubeValid(t *testing.T) {
	var testCases = []struct {
		input    int
		expected string
	}{
		{0, "Special Service"},
		{1, "Closed"},
		{2, "No Service"},
		{3, "No Service"},
		{4, "Planned Closure"},
		{5, "Part Closure"},
		{6, "Severe Delays"},
		{7, "Reduced Service"},
		{8, "Bus Service"},
		{9, "Minor Delays"},
		{10, "Good Service"},
		{11, "Part Closed"},
		{12, "Exit only"},
		{13, "No Step Free Access"},
		{14, "Change of Frequency"},
		{15, "Diverted"},
		{16, "Not Running"},
		{17, "Issues Reported"},
		{18, "No Issues"},
		{19, "Information"},
		{20, "Service Closed"},
	}

	for _, test := range testCases {
		if item, _ := DecodeStatus(test.input, "tube"); item != test.expected {
			t.Error("Expected: ", test.expected, "But got: ", item)
		}
	}
}

func TestTubeInvalid(t *testing.T) {
	var testCases = []struct {
		input int
	}{
		{-2},
		{22},
		{23},
		{12021},
		{12023213},
	}

	for _, test := range testCases {
		if item, err := DecodeStatus(test.input, "tube"); err == nil {
			t.Error("Expected and error, but got: ", "But got: ", item)
		}
	}
}

func TestInvalidSwitches(t *testing.T) {
	var testCases = []struct {
		input string
	}{
		{"someCrap"},
		{"TooB"},
		{"peddleboat"},
		{"riverCanoe"},
		{"tflfail"},
	}

	for _, test := range testCases {
		if item, err := DecodeStatus(0, test.input); err == nil {
			t.Error("Expected and error, but got: ", "But got: ", item)
		}
	}
}
