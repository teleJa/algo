package slidewindow

import (
	"testing"
)

func TestMinWindow(t *testing.T) {

	type field struct {
		s, t string
	}

	tests := []struct {
		name string
		f    field
	}{
		{
			name: "case1",
			f: field{
				s: "ADOBECODEBANC",
				t: "ABC",
			},
		},
		{
			"case2",
			field{
				s: "aa",
				t: "aa",
			},
		},
		{
			"case3",
			field{
				s: "a",
				t: "a",
			},
		},
		{
			"case4",
			field{
				s: "bba",
				t: "ab",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			t.Log(minWindow(test.f.s, test.f.t))
		})
	}

}
