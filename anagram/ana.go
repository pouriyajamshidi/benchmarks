package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func sortString(str string) string {
	var runeSlice []rune

	for _, value := range str {
		runeSlice = append(runeSlice, value)
	}

	sort.SliceStable(runeSlice, func(i, j int) bool {
		return runeSlice[i] < runeSlice[j]
	})

	return string(runeSlice)
}

func contains(strSlice []string, str string) bool {
	for _, v := range strSlice {
		if v == str {
			return true
		}
	}
	return false
}

func main() {
	theDict := make(map[string][]string)

	file, err := os.Open("wordlist.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	words := bufio.NewScanner(file)

	for words.Scan() {
		word := strings.Trim(strings.Replace(words.Text(), "'", "", -1), "\n")

		if len(word) < 2 {
			continue
		}

		sortedWord := sortString(word)

		if _, keyPresent := theDict[sortedWord]; keyPresent {
			if !contains(theDict[sortedWord], word) {
				theDict[sortedWord] = append(theDict[sortedWord], word)
			}
		} else {
			theDict[sortedWord] = []string{}
		}
	}

	counter := 0
	for k, v := range theDict {
		if len(v) >= 10 {
			counter += 1
			fmt.Printf("[%d] %s has %d words: %v\n", counter, k, len(v), v)
		}
	}
}
