package leetcode

import (
	"log"
	"regexp"
	"strings"
	"testing"
)

func TestMoveZeros(t *testing.T) {

	arr := []int{0, 1, 0, 3, 12}
	moveZeroes(arr)
	log.Println(arr)
}

func TestFilterLetterAndNumber(t *testing.T) {

	compile, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		t.Error("regex error:", err)
		return
	}

	allString := compile.ReplaceAllString("A man, a plan, a canal: Panama", "")
	t.Log(allString)
	s := "aA"

	t.Log(strings.EqualFold(string(s[0]), string(s[1])))

}
