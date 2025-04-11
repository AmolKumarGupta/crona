package parser

import (
	"errors"
	"strconv"
	"time"
)

type Flag struct {
	Label string
	Value any
}

type ParseOptions struct {
	Second string
	Minute string
	Hour   string
	Dom    string
	Month  string
	Dow    string
	Flags  []Flag
}

func (p ParseOptions) MatchTime(t time.Time) bool {
	return p.MatchSecond(t) && p.MatchMinute(t) && p.MatchHour(t) && p.MatchDay(t) && p.MatchMonth(t) && p.MatchWeek(t)
}

func (p ParseOptions) MatchSecond(t time.Time) bool {
	if p.Second == "*" {
		return true
	}

	s := t.Second()
	return p.Second == string(rune(s))
}

func (p ParseOptions) MatchMinute(t time.Time) bool {
	if p.Minute == "*" {
		return true
	}

	m := t.Month()
	return p.Minute == string(rune(m))
}

func (p ParseOptions) MatchHour(t time.Time) bool {
	if p.Hour == "*" {
		return true
	}

	h := t.Hour()
	return p.Hour == string(rune(h))
}

func (p ParseOptions) MatchDay(t time.Time) bool {
	if p.Dom == "*" {
		return true
	}

	d := t.Day()
	return p.Dom == string(rune(d))
}

func (p ParseOptions) MatchMonth(t time.Time) bool {
	if p.Month == "*" {
		return true
	}

	m := t.Month()
	return p.Month == string(rune(m))
}

func (p ParseOptions) MatchWeek(t time.Time) bool {
	if p.Dow == "*" {
		return true
	}

	w := t.Weekday()
	return p.Dow == string(rune(w))
}

type bound struct {
	Min   int
	Max   int
	Label map[string]int
}

var (
	SecondBound = bound{Min: 0, Max: 59}
	MinuteBound = bound{Min: 0, Max: 59}
	HourBound   = bound{Min: 0, Max: 23}
	DomBound    = bound{Min: 1, Max: 31}
	MonthBound  = bound{
		Min: 1,
		Max: 12,
		Label: map[string]int{
			"jan": 1,
			"feb": 2,
			"mar": 3,
			"apr": 4,
			"may": 5,
			"jun": 6,
			"jul": 7,
			"aug": 8,
			"sep": 9,
			"oct": 10,
			"nov": 11,
			"dec": 12,
		},
	}
	DowBound = bound{
		Min: 0,
		Max: 6,
		Label: map[string]int{
			"sun": 0,
			"mon": 1,
			"tue": 2,
			"wed": 3,
			"thu": 4,
			"fri": 5,
			"sat": 6,
		},
	}
)

func (b bound) Validate(value string) (bool, error) {
	if value == "*" {
		return true, nil
	}

	num, err := strconv.Atoi(value)
	if err != nil {
		return false, err
	}

	if num < b.Min || num > b.Max {
		return false, errors.New("value out of range")
	}

	return true, nil
}
