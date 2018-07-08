package http

import "testing"

func TestIsASCII(t *testing.T) {
	cases := []struct {
		name string
		in   string
		out  bool
	}{
		{
			name: "is_ascii",
			in:   "abcdefg",
			out:  true,
		},
		{
			name: "is_not_ascii",
			in:   "あいうえお",
			out:  false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if isASCII(c.in) != c.out {
				t.Errorf("string %q should be %v", c.in, c.out)
			}
		})
	}
}

func TestHexEscapeNonASCII(t *testing.T) {
	cases := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "ascii only",
			in:   "abcde",
			out:  "abcde",
		},
		{
			name: "include non ascii",
			in:   "aあい",
			out:  "a%e3%81%82%e3%81%84",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := hexEscapeNonASCII(c.in); got != c.out {
				t.Errorf("want: %v, got: %v", c.out, got)
			}
		})
	}
}
