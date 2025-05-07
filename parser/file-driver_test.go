package parser

import "testing"

func TestIsComment(t *testing.T) {
	tests := []struct {
		line string
		want bool
	}{
		{"// This is a comment", true},
		{"//This is a comment", true},
		{"//", true},
		{"# Another comment", false},
		{"This is not a comment", false},
		{"", false},
	}

	fd := &FileDriver{}

	for _, test := range tests {
		t.Run("isComment", func(t *testing.T) {
			got := fd.isComment(test.line)

			if got != test.want {
				t.Errorf("fd.isComment('%s') = %v; want %v", test.line, got, test.want)
			}
		})
	}
}

func TestSplit(t *testing.T) {
	tests := []struct {
		text string
		want []string
	}{
		{
			"* * * * * * php main.php\n*/5 * * * * * ./sample.sh",
			[]string{"* * * * * * php main.php", "*/5 * * * * * ./sample.sh"},
		},
		{
			"",
			[]string{},
		},
		{
			"* * * * * * php main.php\n\n\n",
			[]string{"* * * * * * php main.php"},
		},
		{
			"* * * * * * php main.php\n\n\n*/5 * * * * * ./sample.sh",
			[]string{"* * * * * * php main.php", "*/5 * * * * * ./sample.sh"},
		},
	}

	for _, test := range tests {
		t.Run("split", func(t *testing.T) {
			fd := &FileDriver{test.text}
			got := fd.split(test.text)

			if len(got) != len(test.want) {
				t.Errorf("fd.split('%s') = %v; want %v", test.text, got, test.want)
				return
			}

			for i, v := range got {
				if v != test.want[i] {
					t.Errorf("fd.split('%s')[%d] = '%s'; want '%s'", test.text, i, v, test.want[i])
				}
			}
		})
	}
}
