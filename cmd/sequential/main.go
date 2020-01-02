package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"

	"github.com/mikehelmick/bogglego/pkg/board"
	"github.com/mikehelmick/bogglego/pkg/trie"
)

// Simple error check
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func loadDictionary(filename string, dict *trie.Trie) {
	fmt.Printf("Loading dictionry from %s\n", filename)
	dat, err := os.Open(filename)
	check(err)
	defer dat.Close()

	scanner := bufio.NewScanner(dat)
	scanner.Split(bufio.ScanLines)

	var wordCount int32 = 0
	for scanner.Scan() {
		dict.AddWord(scanner.Text())
		wordCount++
		if wordCount%1000 == 0 {
			fmt.Print(".")
		}
	}
	fmt.Printf("Loaded %v words.\n", wordCount)
}

func search(prefix string, x, y int, b *board.Board, dict *trie.Trie, visit []bool, words map[string]int) {
	word, err := b.GetAt(x, y)
	if err != nil {
		return
	}
	if visit[x*b.Width()+y] {
		return
	}

	newWord := prefix + word
	if dict.IsWord(newWord) {
		if val, ok := words[newWord]; ok {
			words[newWord] = val + 1
		} else {
			words[newWord] = 1
		}
	}
	if dict.IsPrefix(newWord) {
		visit[x*b.Width()+y] = true
		search(newWord, x-1, y, b, dict, visit, words)
		search(newWord, x+1, y, b, dict, visit, words)
		search(newWord, x, y-1, b, dict, visit, words)
		search(newWord, x, y+1, b, dict, visit, words)
		visit[x*b.Width()+y] = false
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	args := os.Args[1:]
	if len(args) != 1 {
		panic("You must provide a filename to load.")
	}
	dict := trie.New()
	loadDictionary(args[0], dict)

	b := board.New()
	b.PrintBoard()

	fmt.Println("Solving")
	start := time.Now()

	words := make(map[string]int)
	visit := make([]bool, b.Height()*b.Width())

	for x := 0; x < b.Height(); x++ {
		for y := 0; y < b.Width(); y++ {
			search("", x, y, b, dict, visit, words)
		}
	}

	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Printf("Solving took %v\n", elapsed)

	var keys []string
	for k := range words {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	fmt.Printf("Found %v words\n", len(keys))
	for _, k := range keys {
		fmt.Printf("%s : %v\n", k, words[k])
	}
}
