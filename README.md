# Benchmarks

These days, almost all benchmarks out there are extremely unrealistic (check their source code) and not representative of real world applications. So I decided to write my own benchmarks to find an answer suitable for me.

This repo hosts some of the real world applications that I write in different languages to compare and pick the right tool for the task at hand. I also include `developer time` in the benchmarks to give a rough idea and shed some light on **machine** vs **developer** time.

I try to make all the programs as close as possible to each other, not giving an unfair advantage to a language with a certain feature. If I really feel the need to use a certain feature (say threading etc...), I will add it as a separate program.

Please note that the main goal behind this project is to **learn** to solve a problem in different ways and languages to broaden my perspective/skillset and definitely not to put a good/bad label on a language.

Romanticizing a programming language (or a specific technology) is like falling in love with a hammer and totally ignoring the functionality of a screwdriver. Sounds ridiculous, right?

## Available Benchmarks

Here is the list of current benchmarks:

1. [Anagram](anagram_go_nim_vs_python/README.md) comparing `Python`, `Rust`, `Go` and `Nim`.
2. IPTABLES analysis on six million+ lines of log comparing `Python`, `Rust`, `Go` and `Nim` (WIP).

## Contribution

All contributions are welcome. However, please keep in mind that your implementation must be producible in different languages and generic unless we specifically agree otherwise.

All the versions should be kept similar at all times to keep this fair.
