package parser

import (
	"testing"
	"time"
)

func TestMatchSecond(t *testing.T) {
	var tests = []struct {
		optVal string
		input  int
		want   bool
	}{
		{"30", 30, true},
		{"*", 30, true},
		{"15", 30, false},
		{"*/2", 30, true},
	}

	opt := &ParseOptions{}

	for _, test := range tests {
		t.Run(test.optVal, func(t *testing.T) {
			opt.Second = test.optVal

			got := opt.MatchSecond(time.Date(2023, 10, 1, 12, 0, test.input, 0, time.UTC))

			if got != test.want {
				t.Errorf("opt.MatchSecond(%d) = %v; want %v", test.input, got, test.want)
			}
		})
	}
}
