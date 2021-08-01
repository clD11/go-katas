package hackerrank

import (
	"github.com/clD11/go-katas/src/testsupport"
	"testing"
)

func TestBeautifulBinaryString_NONE(t *testing.T) {
	input := "1111111"
	actual := BeautifulBinaryString(input)
	testsupport.AssertThatInt(t, actual, 0)
}

func TestBeautifulBinaryString_ONE(t *testing.T) {
	input := "0101111"
	actual := BeautifulBinaryString(input)
	testsupport.AssertThatInt(t, actual, 1)
}

func TestBeautifulBinaryString_TWO(t *testing.T) {
	input := "0101010"
	actual := BeautifulBinaryString(input)
	testsupport.AssertThatInt(t, actual, 2)
}

func TestBeautifulBinaryString_MIDDLE(t *testing.T) {
	input := "01101011001101111"
	actual := BeautifulBinaryString(input)
	testsupport.AssertThatInt(t, actual, 1)
}
