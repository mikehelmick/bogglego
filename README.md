# bogglego

![Build Status](https://travis-ci.com/mikehelmick/bogglego.svg?branch=master)

A Boggle solver written in Go. Uses a Trie to represent the dictionary.
There are two versions, a sequential implementation and parallel (from each
spot).

# To Run

Use the scripts directory

```
./scripts/sequential.sh
./scripts/parallel.sh
```

# Run tests

```
go test ./... -count 1
go test ./... --bench . -count 2
```
