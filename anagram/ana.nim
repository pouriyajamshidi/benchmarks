import std/[tables, strutils, algorithm, strformat]


var theDict = initTable[string, seq[string]]()

for word in lines("wordlist.txt"):
    var word = word.replace("'", "").strip()

    if len(word) < 2:
        continue

    var sortedWord = join(sorted(word), "")

    if theDict.hasKey(sortedWord):
        if not theDict[sortedWord].contains(word):
            theDict[sortedWord].add(word)
    else:
        theDict[sortedWord] = @[]


var counter = 0
for k, v in theDict:
    if len(v) >= 10:
        counter += 1
        echo fmt"[{counter}] {k} has {len(v)} words: {v}"
