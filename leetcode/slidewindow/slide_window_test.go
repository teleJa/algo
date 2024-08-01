package slidewindow

import (
	"testing"
)

func TestMinWindow(t *testing.T) {
	s := "ADOBECODEBANC"
	su := "ABC"
	t.Log(minWindow(s, su) == "BANC")

	s = "aa"
	su = "aa"
	t.Log(minWindow(s, su) == "aa")

}
