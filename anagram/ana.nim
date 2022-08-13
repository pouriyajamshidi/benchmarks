import std/[tables, strutils, algorithm, strformat]


var theDict = initTable[string, seq[string]]()

for word in lines(("wordlist.txt")):
    var word = word.replace("'", "").strip()
    if len(word) < 2:
        continue

    var sorted_word = join(sorted(word), "")

    if theDict.hasKey(sorted_word):
        if not theDict[sorted_word].contains(word):
            theDict[sorted_word].add(word)
    else:
        theDict[sorted_word] = @[]


var counter = 0
for k, v in theDict:
    if len(v) >= 10:
        counter += 1
        echo fmt"[{counter}] {k} has {len(v)} words: {v}"
