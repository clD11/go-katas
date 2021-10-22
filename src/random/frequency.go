package random

import (
	"bufio"
	"io"
	"os"
	"sort"
	"strings"
)

type Blog struct {
	title string
	like  bool
}

func Frequency(input []Blog, n int) []string {
	freq := make(map[string]int)
	for _, b := range input {
		if b.like {
			freq[b.title] += 1
		}
	}

	rev := make(map[int]string)
	for k, v := range freq {
		rev[v] = k
	}

	s := make([]int, 0)
	for k, _ := range rev {
		s = append(s, k)
	}
	sort.Ints(s)

	var r = make([]string, 0, n)
	for i := len(s) - 1; i > -1 && n > 0; i-- {
		r = append(r, rev[s[i]])
		n--
	}

	return r
}

func GetPlayerFrequency(filename string, players int) []string {
	f, err := os.Open(filename)
	check(err)
	var playerToFrequency = make(map[string]int)
	var reader = bufio.NewReader(f)
	for {
		name, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		playerToFrequency[name] += 1
	}

	var scores = make(map[int]string)
	for k, v := range playerToFrequency {
		scores[v] = k
	}

	keys := make([]int, 0, len(scores))
	for k := range scores {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	var result = make([]string, 0, players)
	for i := len(keys) - 1; i > -1 && players > 0; i-- {
		s := strings.TrimRight(scores[keys[i]], "\r\n")
		result = append(result, s)
		players--
	}

	return result
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
