#!/usr/bin/env python3


with open("wordlist.txt", "r", errors="ignore") as f:
    words = f.readlines()


the_dict: dict = {}

for word in words:
    word = word.replace("'", "").strip()
    if len(word) < 2:
        continue

    sorted_word = "".join(sorted(word))

    if sorted_word in the_dict:
        if word not in the_dict[sorted_word]:
            the_dict[sorted_word].append(word)
    else:
        the_dict[sorted_word] = []

counter = 0
for k, v in the_dict.items():
    if len(v) >= 10:
        counter += 1
        print(f"[{counter}] {k} has {len(v)} words: {v}")
