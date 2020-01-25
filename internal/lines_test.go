package api

import (
	"testing"
)

var testCases = []Lines{
	{"bakerloo", "Bakerloo", "tube", "", testCasesLineStatus},
	{"central", "Central", "tube", "", testCasesLineStatus},
	{"circle", "Circle", "tube", "", testCasesLineStatus},
	{"district", "District", "tube", "", testCasesLineStatus},
	{"hammersmith-city", "Hammersmith & City", "tube", "", testCasesLineStatus},
	{"jubilee", "Jubilee", "tube", "", testCasesLineStatus},
	{"metropolitan", "Metropolitan", "tube", "", testCasesLineStatus},
	{"northern", "Northern", "tube", "", testCasesLineStatus},
	{"piccadilly", "Piccadilly", "tube", "", testCasesLineStatus},
	{"victoria", "Victoria", "tube", "", testCasesLineStatus},
	{"waterloo-city", "Waterloo & City", "tube", "", testCasesLineStatus},
}

var testCasesLineStatus = []lineStatus{
	{0, "", 10, "Good Service", "", []validityPeriods{}},
	{0, "Metropolitan", 5, "Part Closure", "METROPOLITAN LINE: Saturday 25 and Sunday 26 January, no service between Wembley Park and Aldgate, please use Jubilee line services, or Chiltern Railways between Harrow-on-the-Hill and Marylebone.", []validityPeriods{}},
}

func TestDetermineStatus(t *testing.T) {
	var testCases = []struct {
		input    testCases
		expected []LineSummary
	}{
		{testCases[0], []LineSummary{}},
	}

	for _, test := range testCases {
		if item, _ := DetermineStatus(test.input, test.expected); item != test.expected {
			t.Error("Expected: ", test.expected, "But got: ", item)
		}
	}
}
