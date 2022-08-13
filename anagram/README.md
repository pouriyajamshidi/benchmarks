# Anagram

This is a Kata that I came across a few months ago. Though it is not truly a real world problem, it definitely embodies some of my daily challenges.

I have tried to make all the versions as close as possible.

Here is the link to the [original Kata](http://codekata.com/kata/kata06-anagrams/).

## Table of Contents

- [Anagram](#anagram)
  - [Table of Contents](#table-of-contents)
  - [Test bench](#test-bench)
  - [Languages](#languages)
  - [Prerequisites](#prerequisites)
    - [Python](#python)
    - [Nim](#nim)
    - [Go](#go)
  - [Run](#run)
  - [Results](#results)

## Test bench

| OS                 | Kernel     | CPU                            | Memory |
| ------------------ | ---------- | ------------------------------ | ------ |
| Ubuntu 20.04.4 LTS | 5.10.102.1 | Intel i7-7700HQ (8) @ 2.807GHz | 24GB   |

## Languages

| Language | Version  |
| -------- | -------- |
| Python   | `3.10.5` |
| Nim      | `1.7.1`  |
| Go       | `1.19`   |

## Prerequisites

Make sure your versions match the [test bench](#test-bench)
Benchmarking tool: [Hyperfine](https://github.com/sharkdp/hyperfine)

### Python

Check Python version:

```bash
$ python -V
Python 3.10.5
```

### Nim

Check and compile Nim version:

```bash
$ nim -V
Nim Compiler Version 1.7.1 [Linux: amd64]
Compiled at 2022-08-12
Copyright (c) 2006-2022 by Andreas Rumpf

git hash: ff25103c9ab9d51821e9e8641955c8d24f7db6b8
```

```bash
nim c -d:release -mm:orc --threads:off -o:ananim ana.nim
```

> with or without threads, the result is the same

### Go

Check and compile Go version:

```bash
$ go version
go version go1.19 linux/amd64
```

```bash
go build -o anago ana.go
```

Please `unrar` the `wordlist.rar` before running the benchmark.

```bash
unrar x wordlist.rar
```

## Run

Drop the current caches:

```bash
sync; echo 3 | sudo tee /proc/sys/vm/drop_caches
```

Run the benchmark:

```bash
hyperfine "python3.10 ana.py" "./ananim" "./anago"
```

## Results

Go managed to come out first but oddly enough, Nim is the slowest.

| Command  |     Mean [ms] | Min [ms] | Max [ms] |    Relative |
| :------- | ------------: | -------: | -------: | ----------: |
| `Go`     |  787.6 ± 18.5 |    762.0 |    820.2 |        1.00 |
| `Python` |  874.7 ± 75.2 |    785.8 |   1047.2 | 1.11 ± 0.10 |
| `Nim`    | 1502.0 ± 47.3 |   1405.6 |   1565.1 | 1.91 ± 0.07 |
