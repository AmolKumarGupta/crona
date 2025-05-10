package parser

import (
	"strings"
	"testing"
	"time"
)

func TestMatchSecond(t *testing.T) {
	var tests = []struct {
		optVal string
		input  int
		want   bool
	}{
		{"*", 30, true},

		{"10,0,30", 0, true},
		{"10,20,30", 30, true},
		{"10,20,45", 30, false},
		{"10,20,99", 30, false},
		{"15,20", 15, true},

		{"0-10", 0, true},
		{"0-10", 30, false},
		{"10-0", 30, false},
		{"0-99", 30, false},

		{"*/2", 30, true},
		{"*/10", 30, true},
		{"*/30", 30, true},
		{"*/59", 30, false},
		{"*/60", 30, false},

		{"0", 0, true},
		{"59", 59, true},
		{"15", 30, false},
		{"99", 30, false},
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

func TestMatchMinute(t *testing.T) {
	var tests = []struct {
		optVal string
		input  int
		want   bool
	}{
		{"*", 30, true},

		{"10,0,30", 0, true},
		{"10,20,30", 30, true},
		{"10,20,45", 30, false},
		{"10,20,99", 30, false},

		{"0-10", 0, true},
		{"0-10", 30, false},
		{"10-0", 30, false},
		{"0-99", 30, false},

		{"*/2", 30, true},
		{"*/10", 30, true},
		{"*/30", 30, true},
		{"*/59", 30, false},
		{"*/60", 30, false},

		{"0", 0, true},
		{"59", 59, true},
		{"15", 30, false},
		{"99", 30, false},
	}

	opt := &ParseOptions{}

	for _, test := range tests {
		t.Run(test.optVal, func(t *testing.T) {
			opt.Minute = test.optVal

			got := opt.MatchMinute(time.Date(2023, 10, 1, 12, test.input, 0, 0, time.UTC))

			if got != test.want {
				t.Errorf("opt.MatchMinute(%d) = %v; want %v", test.input, got, test.want)
			}
		})
	}
}

func TestMatchHour(t *testing.T) {
	var tests = []struct {
		optVal string
		input  int
		want   bool
	}{
		{"*", 0, true},

		{"10,0,23", 23, true},
		{"10,20,23", 20, true},
		{"10,20,23", 22, false},
		{"10,20,99", 22, false},

		{"0-10", 0, true},
		{"0-10", 22, false},
		{"10-0", 10, false},
		{"0-99", 10, false},

		{"*/2", 12, true},
		{"*/12", 0, true},
		{"*/2", 21, false},
		{"*/10", 21, false},
		{"*/59", 21, false},
		{"*/60", 21, false},

		{"0", 0, true},
		{"23", 23, true},
		{"15", 16, false},
		{"99", 16, false},
	}

	opt := &ParseOptions{}

	for _, test := range tests {
		t.Run(test.optVal, func(t *testing.T) {
			opt.Hour = test.optVal

			got := opt.MatchHour(time.Date(2023, 10, 1, test.input, 0, 0, 0, time.UTC))

			if got != test.want {
				t.Errorf("opt.MatchHour(%d) = %v; want %v", test.input, got, test.want)
			}
		})
	}
}

func TestMatchDay(t *testing.T) {
	var tests = []struct {
		optVal string
		input  int
		want   bool
	}{
		{"*", 1, true},
		{"*", 31, true},
		{"*", 0, true},

		{"10,1,23", 23, true},
		{"10,20,23", 20, true},
		{"10,20,23", 22, false},
		{"10,20,99", 22, false},

		{"1-10", 0, false},
		{"1-10", 22, false},
		{"10-1", 10, false},
		{"1-99", 10, false},

		{"*/2", 13, true},
		{"*/12", 1, true},
		{"*/2", 21, true},
		{"*/10", 21, true},
		{"*/10", 20, false},
		{"*/59", 21, false},
		{"*/60", 21, false},

		{"0", 0, false},
		{"23", 23, true},
		{"15", 16, false},
		{"99", 16, false},
	}

	opt := &ParseOptions{}

	for _, test := range tests {
		t.Run(test.optVal, func(t *testing.T) {
			opt.Dom = test.optVal

			got := opt.MatchDay(time.Date(2023, 10, test.input, 0, 0, 0, 0, time.UTC))

			if got != test.want {
				t.Errorf("opt.MatchDay(%d) = %v; want %v", test.input, got, test.want)
			}
		})
	}
}

func TestMatchMonth(t *testing.T) {
	var tests = []struct {
		optVal string
		input  int
		want   bool
	}{
		{"*", 1, true},
		{"*", 12, true},

		{"10,1,12", 1, true},
		{"10,20,99", 11, false},

		{"1-10", 1, true},
		{"1-10", 11, false},
		{"10-1", 10, false},
		{"1-99", 10, false},

		{"*/2", 11, true},
		{"*/12", 1, true},
		{"*/2", 3, true},
		{"*/10", 11, true},
		{"*/10", 12, false},
		{"*/59", 1, false},
		{"*/60", 1, false},

		{"0", 0, false},
		{"2", 2, true},
		{"12", 11, false},
		{"99", 11, false},
	}

	opt := &ParseOptions{}

	for _, test := range tests {
		t.Run(test.optVal, func(t *testing.T) {
			opt.Month = test.optVal

			got := opt.MatchMonth(time.Date(2023, time.Month(test.input), 1, 0, 0, 0, 0, time.UTC))

			if got != test.want {
				t.Errorf("opt.MatchMonth(%d) = %v; want %v", test.input, got, test.want)
			}
		})
	}
}

func TestMatchWeek(t *testing.T) {
	var tests = []struct {
		optVal string
		input  int
		want   bool
	}{
		{"*", 0, true},
		{"*", 6, true},

		{"0,1,6", 0, true},
		{"0,1,6", 1, true},
		{"0,1,6", 5, false},
		{"0,1,6", 7, false},

		{"0-3", 0, true},
		{"0-3", 3, true},
		{"0-3", 4, false},
		{"3-0", 3, false},
		{"0-7", 6, false},

		{"*/2", 0, true},
		{"*/2", 2, true},
		{"*/2", 4, true},
		{"*/2", 5, false},
		{"*/7", 6, false},
		{"*/8", 6, false},

		{"0", 0, true},
		{"6", 6, true},
		{"3", 4, false},
		{"7", 6, false},
	}

	opt := &ParseOptions{}

	// base date with known weekday (Sunday == 0)
	base := time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC)

	for _, test := range tests {
		t.Run(test.optVal, func(t *testing.T) {
			opt.Dow = test.optVal
			var got bool

			if test.input < 0 || test.input > 6 {
				got = false

			} else {
				offset := (test.input - int(base.Weekday()) + 7) % 7
				d := base.AddDate(0, 0, offset)
				got = opt.MatchWeek(d)
			}

			if got != test.want {
				t.Errorf("opt.MatchWeek(%d) = %v; want %v", test.input, got, test.want)
			}
		})
	}
}

func TestMatchTime(t *testing.T) {
	var tests = []struct {
		val   []string
		input time.Time
		want  bool
	}{
		{[]string{"*", "*", "*", "*", "*", "*"}, time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC), true},
		{[]string{"0", "0", "9-17", "*", "*", "*"}, time.Date(2023, 10, 1, 10, 0, 0, 0, time.UTC), true},
		{[]string{"0", "0", "9-17", "*", "*", "*"}, time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC), false},
		{[]string{"0", "0", "0", "*", "*", "*"}, time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC), true},
		{[]string{"0", "0", "0", "*", "*", "*"}, time.Date(2023, 10, 2, 0, 0, 0, 0, time.UTC), true},
		{[]string{"0", "0", "0", "*", "*", "0"}, time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC), true},
		{[]string{"0", "0", "0", "*", "*", "0"}, time.Date(2023, 10, 8, 0, 0, 0, 0, time.UTC), true},
		{[]string{"0", "0", "0", "1", "*/3", "*"}, time.Date(2023, 9, 1, 0, 0, 0, 0, time.UTC), false},
		{[]string{"0", "0", "0", "1", "*/3", "*"}, time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC), true},
		{[]string{"0", "0", "0", "1", "*/3", "*"}, time.Date(2023, 10, 2, 0, 0, 0, 0, time.UTC), false},
	}

	opt := &ParseOptions{}

	for _, test := range tests {
		t.Run(strings.Join(test.val, " "), func(t *testing.T) {
			opt.Second = test.val[0]
			opt.Minute = test.val[1]
			opt.Hour = test.val[2]
			opt.Dom = test.val[3]
			opt.Month = test.val[4]
			opt.Dow = test.val[5]

			got := opt.MatchTime(test.input)

			if got != test.want {
				t.Errorf("opt.MatchTime(%v) = %v; want %v", test.input, got, test.want)
			}
		})
	}
}

func TestValidate(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"*", true},
		{"1-5", true},
		{"1,5", true},
		{"*/3", true},
		{"abc", false},
		{"1000", false},
		{"1_5", false},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			got, _ := SecondBound.Validate(test.input)

			if got != test.want {
				t.Errorf("SecondBound.Validate(%s) = %v; want %v", test.input, got, test.want)
			}
		})
	}
}

func TestCompare(t *testing.T) {
	po1 := NewParseOptions("*", "*", "*", "*", "*", "*", []Flag{{"--test", ""}, {"--d", "info"}})
	po2 := NewParseOptions("*", "*", "*", "*", "*", "*", []Flag{{"--test", ""}, {"--d", "info"}})
	po3 := NewParseOptions("*", "*", "*", "*", "*", "*", []Flag{{"--test", ""}, {"--d", "debug"}})
	po4 := NewParseOptions("*", "*", "*", "*", "*", "*", []Flag{{"--n", ""}, {"--d", "debug"}})

	if !po1.Compare(*po2) {
		t.Error("po1 should be same as po2", po1, po2)
	}

	if po1.Compare(*po3) {
		t.Error("po1 should be not same as po3", po1, po3)
	}

	if po1.Compare(*po4) {
		t.Error("po1 should be not same as po4", po1, po4)
	}
}
