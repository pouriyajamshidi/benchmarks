module main

import os


fn main() {
	mut the_dict := map[string][]string{}
	mut words := os.read_lines("wordlist.txt")?

	for word in words {
		mut word_mod := word.replace("'", "").trim("\n")

		if word_mod.len < 2 {
			continue
		}

		mut unsorted_word := word_mod.split("")

		unsorted_word.sort_with_compare(fn (a &string, b &string) int {
				if a < b {
					return -1
				}
				if a > b {
					return 1
				}
				return 0
			}
		)

		mut sorted_word := unsorted_word.join("")

		if sorted_word in the_dict {
			if !(word_mod in the_dict[sorted_word]) {
				the_dict[sorted_word] << word_mod
			}
		} else {
			the_dict[sorted_word] = []string{}
		}
	}

	mut counter := 0
	for k, v in the_dict {
		if v.len >= 10 {
			counter += 1
			println('$counter $k has $v.len words: $v')
		}
	}
}
