package api

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

var testCasesLineStatus = []lineStatus{
	{0, "", 10, "Good Service", "", createdTime{}, []validityPeriods{validityPeriods{From: fromTime{}, To: toTime{}, IsNow: true}}},
	{0, "Metropolitan", 5, "Part Closure", "METROPOLITAN LINE: Saturday 25 and Sunday 26 January, no service between Wembley Park and Aldgate, please use Jubilee line services, or Chiltern Railways between Harrow-on-the-Hill and Marylebone.", createdTime{}, []validityPeriods{validityPeriods{From: fromTime{}, To: toTime{}, IsNow: true}}},
}

func TestDetermineStatus(t *testing.T) {
	var testCases = []struct {
		input    []Lines
		expected []LineSummary
	}{
		{[]Lines{Lines{"bakerloo", "Bakerloo", "tube", "", time.Time{}, time.Time{}, testCasesLineStatus}}, []LineSummary{LineSummary{"Bakerloo", "Good Service", ""}}},
		{[]Lines{Lines{"central", "Central", "tube", "", time.Time{}, time.Time{}, testCasesLineStatus}}, []LineSummary{LineSummary{"Central", "Good Service", ""}}},
		{[]Lines{Lines{"circle", "Circle", "tube", "", time.Time{}, time.Time{}, testCasesLineStatus}}, []LineSummary{LineSummary{"Circle", "Good Service", ""}}},
		{[]Lines{Lines{"district", "District", "tube", "", time.Time{}, time.Time{}, testCasesLineStatus}}, []LineSummary{LineSummary{"District", "Good Service", ""}}},
		{[]Lines{Lines{"hammersmith-city", "Hammersmith & City", "tube", "", time.Time{}, time.Time{}, testCasesLineStatus}}, []LineSummary{LineSummary{"Hammersmith & City", "Good Service", ""}}},
		{[]Lines{Lines{"jubilee", "Jubilee", "tube", "", time.Time{}, time.Time{}, testCasesLineStatus}}, []LineSummary{LineSummary{"Jubilee", "Good Service", ""}}},
		{[]Lines{Lines{"metropolitan", "Metropolitan", "tube", "", time.Time{}, time.Time{}, testCasesLineStatus}}, []LineSummary{LineSummary{"Metropolitan", "Good Service", ""}}},
		{[]Lines{Lines{"northern", "Northern", "tube", "", time.Time{}, time.Time{}, testCasesLineStatus}}, []LineSummary{LineSummary{"Northern", "Good Service", ""}}},
		{[]Lines{Lines{"piccadilly", "Piccadilly", "tube", "", time.Time{}, time.Time{}, testCasesLineStatus}}, []LineSummary{LineSummary{"Piccadilly", "Good Service", ""}}},
		{[]Lines{Lines{"victoria", "Victoria", "tube", "", time.Time{}, time.Time{}, testCasesLineStatus}}, []LineSummary{LineSummary{"Victoria", "Good Service", ""}}},
		{[]Lines{Lines{"waterloo-city", "Waterloo & City", "tube", "", time.Time{}, time.Time{}, testCasesLineStatus}}, []LineSummary{LineSummary{"Waterloo & City", "Good Service", ""}}},
	}

	for _, test := range testCases {

		output := determineStatus(test.input)

		if !cmp.Equal(output[0], test.expected[0]) {
			t.Error("Expected: ", test.expected[0], "But got: ", output[0])
		}
	}
}
