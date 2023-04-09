package yr_test

import (
	"testing"

	"github.com/yildirimsinop/minyr/yr"
)

func TestCountLines(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: "kjevik-temp-celsius-20220318-20230318.csv", want: 25},
	}

	for _, tc := range tests {
		got := yr.CountLines(tc.input)
		if got != tc.want {
			t.Errorf("%v: want %v, got %v,", tc.input, tc.want, got)
		}
	}
}
func TestGetAverageTemperature(t *testing.T) {
	actualAvg, err := yr.GetAverageTemperature("kjevik-temp-celsius-20220318-20230318.csv", "celsius")
	if err != nil {
		t.Fatal(err)
	}

	expectedAvg := "-0.60"
	if actualAvg != expectedAvg {
		t.Errorf("expected average temperature %v, but got %v", expectedAvg, actualAvg)
	}
}
