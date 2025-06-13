package parser_test

import (
	"testing"

	"github.com/AmolKumarGupta/crona/parser"
)

func TestGet(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"fs", true},
		{"mem", true},
		{"none", false},
	}

	for _, test := range tests {
		t.Run("Get", func(t *testing.T) {
			driver, _ := parser.Get(test.input)

			if test.want && driver == nil {
				t.Errorf("Get('%s') = %v, want %v", test.input, driver, test.want)
			}

			if !test.want && driver != nil {
				t.Errorf("Get('%s') = %v, want %v", test.input, driver, test.want)
			}
		})
	}
}
