package api

import (
	"testing"
)

func TestDetermineStatus(t *testing.T) {
	var testCases = []struct {
		input    []Lines
		expected []LineSummary
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
		if item, _ := DetermineStatus(test.input, "cableCar"); item != test.expected {
			t.Error("Expected: ", test.expected, "But got: ", item)
		}
	}
}
