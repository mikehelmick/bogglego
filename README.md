# bogglego

[![Build Status](https://travis-ci.com/mikehelmick/bogglego.svg?branch=master)](https://travis-ci.com/mikehelmick/bogglego)
[![Report Card](https://goreportcard.com/badge/github.com/mikehelmick/bogglego)](https://goreportcard.com/report/github.com/mikehelmick/bogglego)

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
