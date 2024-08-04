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

func Test_checkInclusion(t *testing.T) {

	type field struct {
		s1, s2 string
	}

	tests := []struct {
		name string
		f    field
	}{
		{
			"case1",
			field{
				"ab",
				"eidbaooo",
			},
		},
		{
			"case2",
			field{
				s1: "ab",
				s2: "eidboaoo",
			},
		},
		{
			"case3",
			field{
				s1: "hello",
				s2: "ooolleoooleh",
			},
		},
		{
			"case4",
			field{
				s1: "abc",
				s2: "ccccbbbbaaaa",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Log(checkInclusion(test.f.s1, test.f.s2))
		})
	}

}
