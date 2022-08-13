package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func sortString(s string) string {
	var r []rune

	for _, value := range s {
		r = append(r, value)
	}

	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})

	return string(r)
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func main() {
	file, err := os.Open("wordlist.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	theDict := make(map[string][]string)

	words := bufio.NewScanner(file)

	for words.Scan() {
		word := words.Text()
		word = strings.Trim(strings.Replace(word, "'", "", -1), "\n")

		if len(word) < 2 {
			continue
		}

		sortedWord := sortString(word)

		if theDict[sortedWord] != nil {
			if !contains(theDict[sortedWord], word) {
				theDict[sortedWord] = append(theDict[sortedWord], word)
			}
		} else {
			theDict[sortedWord] = []string{word}
		}
	}

	counter := 0
	for k, v := range theDict {
		if len(v) > 10 {
			counter += 1
			fmt.Printf("[%d] %s has %d words: %v\n", counter, k, len(v), v)
		}
	}
}
