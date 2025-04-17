package parser

import (
	"errors"
	"slices"
	"strconv"
	"strings"
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

func (p ParseOptions) MatchSecond(t time.Time) bool {
	return p.matchTimeValue(t.Second(), p.Second, SecondBound)
}

func (p ParseOptions) MatchMinute(t time.Time) bool {
	return p.matchTimeValue(t.Minute(), p.Minute, MinuteBound)
}

func (p ParseOptions) MatchHour(t time.Time) bool {
	return p.matchTimeValue(t.Hour(), p.Hour, HourBound)
}

func (p ParseOptions) MatchDay(t time.Time) bool {
	return p.matchTimeValue(t.Day(), p.Dom, DomBound)
}

func (p ParseOptions) MatchMonth(t time.Time) bool {
	return p.matchTimeValue(int(t.Month()), p.Month, MonthBound)
}

func (p ParseOptions) MatchWeek(t time.Time) bool {
	return p.matchTimeValue(int(t.Weekday()), p.Dow, DowBound)
}

func (p ParseOptions) MatchTime(t time.Time) bool {
	return p.MatchSecond(t) && p.MatchMinute(t) && p.MatchHour(t) && p.MatchDay(t) && p.MatchMonth(t) && p.MatchWeek(t)
}

func (p ParseOptions) matchTimeValue(val int, str string, bnd bound) bool {
	if isAllUnit(str) {
		return true
	}

	if isMultipleValues(str, bnd) {
		items := strings.Split(str, ",")
		return slices.Contains(items, strconv.Itoa(val))
	}

	if isRange(str, bnd) {
		nums := strings.Split(str, "-")
		start, err := strconv.Atoi(nums[0])
		if err != nil {
			return false
		}

		end, err := strconv.Atoi(nums[1])
		if err != nil {
			return false
		}
		return val >= start && val <= end
	}

	if isStepRange(str, bnd) {
		str := strings.TrimPrefix(str, "*/")
		num, err := strconv.Atoi(str)
		if err != nil {
			return false
		}

		offset := bnd.Min
		ranges := []int{}

		for i := offset; i <= bnd.Max; i += num {
			ranges = append(ranges, i)
		}

		return slices.Contains(ranges, val)
	}

	s := val
	d, err := strconv.Atoi(str)
	if err != nil {
		return false
	}

	return d == s
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
	if isAllUnit(value) {
		return true, nil
	}

	if isMultipleValues(value, b) {
		return true, nil
	}

	if isRange(value, b) {
		return true, nil
	}

	if isStepRange(value, b) {
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

// isAllUnit checks if the string is a wildcard "*"
// indicating that all values are accepted.
func isAllUnit(str string) bool {
	return str == "*"
}

// isMultipleValues checks if the string contains multiple values separated by commas.
// It also checks if each value is within the specified bounds.
func isMultipleValues(str string, bnd bound) bool {
	items := strings.Split(str, ",")
	if !(len(items) > 1) {
		return false
	}

	for _, item := range items {
		num, err := strconv.Atoi(item)

		if err != nil {
			return false
		}

		if num < bnd.Min || num > bnd.Max {
			return false
		}
	}

	return true
}

// isRange checks if the string is in the format "N-M"
// where N and M are numbers within the specified bounds.
func isRange(str string, bnd bound) bool {
	if !strings.Contains(str, "-") {
		return false
	}

	nums := strings.Split(str, "-")

	if len(nums) != 2 {
		return false
	}

	start, err := strconv.Atoi(nums[0])
	if err != nil {
		return false
	}
	end, err := strconv.Atoi(nums[1])
	if err != nil {
		return false
	}

	return start >= bnd.Min && end <= bnd.Max && start <= end
}

// isStepRange checks if the string is in the format "*/N"
// where N is a number within the specified bounds.
func isStepRange(str string, bnd bound) bool {
	if !strings.HasPrefix(str, "*/") {
		return false
	}

	str = strings.TrimPrefix(str, "*/")

	num, err := strconv.Atoi(str)
	if err != nil {
		return false
	}

	return num >= bnd.Min && num <= bnd.Max
}
