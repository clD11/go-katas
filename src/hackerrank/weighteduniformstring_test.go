package hackerrank

import (
	"bufio"
	"github.com/clD11/go-katas/src/testsupport"
	"log"
	"os"
	"strconv"
	"testing"
)

func TestWeightedUniformStrings(t *testing.T) {
	s := "abccddde"
	queries := []int32{1, 3, 12, 5, 9, 10}
	actual := WeightedUniformStrings(s, queries)
	expected := []string{"Yes", "Yes", "Yes", "Yes", "No", "No"}
	for i := 0; i < len(queries); i++ {
		testsupport.AssertThatString(t, actual[i], expected[i])
	}
}

func TestWeightedUniformStrings_One(t *testing.T) {
	s := "a"
	queries := []int32{1}
	actual := WeightedUniformStrings(s, queries)
	expected := []string{"Yes"}
	for i := 0; i < len(queries); i++ {
		testsupport.AssertThatString(t, actual[i], expected[i])
	}
}

func TestWeightedUniformStrings_Two(t *testing.T) {
	s := "aa"
	queries := []int32{2}
	actual := WeightedUniformStrings(s, queries)
	expected := []string{"Yes"}
	for i := 0; i < len(queries); i++ {
		testsupport.AssertThatString(t, actual[i], expected[i])
	}
}

func TestWeightedUniformStrings_Three(t *testing.T) {
	s := "c"
	queries := []int32{3}
	actual := WeightedUniformStrings(s, queries)
	expected := []string{"Yes"}
	for i := 0; i < len(queries); i++ {
		testsupport.AssertThatString(t, actual[i], expected[i])
	}
}

func TestWeightedUniformStrings_Four(t *testing.T) {
	s := "aab"
	queries := []int32{4}
	actual := WeightedUniformStrings(s, queries)
	expected := []string{"No"}
	for i := 0; i < len(queries); i++ {
		testsupport.AssertThatString(t, actual[i], expected[i])
	}
}

func TestWeightedUniformStrings_Nine(t *testing.T) {
	s := "abccddde"
	queries := []int32{9}
	actual := WeightedUniformStrings(s, queries)
	expected := []string{"No"}
	for i := 0; i < len(queries); i++ {
		testsupport.AssertThatString(t, actual[i], expected[i])
	}
}

func TestWeightedUniformStrings_Twelve(t *testing.T) {
	s := "acccce"
	queries := []int32{12}
	actual := WeightedUniformStrings(s, queries)
	expected := []string{"Yes"}
	for i := 0; i < len(queries); i++ {
		testsupport.AssertThatString(t, actual[i], expected[i])
	}
}

func TestWeightedUniformStrings_A(t *testing.T) {
	s := "aaazzzzaa"
	queries := []int32{12}
	actual := WeightedUniformStrings(s, queries)
	expected := []string{"Yes"}
	for i := 0; i < len(queries); i++ {
		testsupport.AssertThatString(t, actual[i], expected[i])
	}
}

func TestWeightedUniformStrings_Long(t *testing.T) {
	f, _ := os.Open("testdata/hackerrank_weighted_strings.txt")
	defer f.Close()

	s := bufio.NewScanner(f)

	s.Scan()
	l := s.Text()

	s.Scan()
	q, _ := strconv.Atoi(s.Text())

	queries := make([]int32, q, q)
	for i := 0; i < q; i++ {
		s.Scan()
		num, _ := strconv.Atoi(s.Text())
		queries[i] = int32(num)
	}

	f, _ = os.Open("testdata/hackerrank_weighted_strings_ans.txt")
	defer f.Close()

	s = bufio.NewScanner(f)

	rrr := WeightedUniformStrings(l, []int32{389400})
	log.Print(rrr)

	res := WeightedUniformStrings(l, queries)

	for i := 0; i < q; i++ {
		s.Scan()
		r := s.Text()
		if r != res[i] {
			log.Printf("exp %s act %s query %d", r, res[i], queries[i])
		}
	}
}
