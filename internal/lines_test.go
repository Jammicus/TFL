package api

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

var testCasesLineStatus = []lineStatus{
	{0, "", 10, "Good Service", "", []validityPeriods{}},
	{0, "Metropolitan", 5, "Part Closure", "METROPOLITAN LINE: Saturday 25 and Sunday 26 January, no service between Wembley Park and Aldgate, please use Jubilee line services, or Chiltern Railways between Harrow-on-the-Hill and Marylebone.", []validityPeriods{}},
}

func TestDetermineStatus(t *testing.T) {
	var testCases = []struct {
		input    []Lines
		expected []LineSummary
	}{
		{[]Lines{Lines{"bakerloo", "Bakerloo", "tube", "", testCasesLineStatus}}, []LineSummary{LineSummary{"Bakerloo", "Good Service", ""}}},
		{[]Lines{Lines{"central", "Central", "tube", "", testCasesLineStatus}}, []LineSummary{LineSummary{"Central", "Good Service", ""}}},
		{[]Lines{Lines{"circle", "Circle", "tube", "", testCasesLineStatus}}, []LineSummary{LineSummary{"Circle", "Good Service", ""}}},
		{[]Lines{Lines{"district", "District", "tube", "", testCasesLineStatus}}, []LineSummary{LineSummary{"District", "Good Service", ""}}},
		{[]Lines{Lines{"hammersmith-city", "Hammersmith & City", "tube", "", testCasesLineStatus}}, []LineSummary{LineSummary{"Hammersmith & City", "Good Service", ""}}},
		{[]Lines{Lines{"jubilee", "Jubilee", "tube", "", testCasesLineStatus}}, []LineSummary{LineSummary{"Jubilee", "Good Service", ""}}},
		{[]Lines{Lines{"metropolitan", "Metropolitan", "tube", "", testCasesLineStatus}}, []LineSummary{LineSummary{"Metropolitan", "Good Service", ""}}},
		{[]Lines{Lines{"northern", "Northern", "tube", "", testCasesLineStatus}}, []LineSummary{LineSummary{"Northern", "Good Service", ""}}},
		{[]Lines{Lines{"piccadilly", "Piccadilly", "tube", "", testCasesLineStatus}}, []LineSummary{LineSummary{"Piccadilly", "Good Service", ""}}},
		{[]Lines{Lines{"victoria", "Victoria", "tube", "", testCasesLineStatus}}, []LineSummary{LineSummary{"Victoria", "Good Service", ""}}},
		{[]Lines{Lines{"waterloo-city", "Waterloo & City", "tube", "", testCasesLineStatus}}, []LineSummary{LineSummary{"Waterloo & City", "Good Service", ""}}},
	}

	for _, test := range testCases {

		output := DetermineStatus(test.input)

		if !cmp.Equal(output[0], test.expected[0]) {
			t.Error("Expected: ", test.expected[0], "But got: ", output[0])
		}
	}
}
