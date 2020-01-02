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

func search(prefix string, x, y int, b *board.Board, dict *trie.Trie, visit []bool, ch chan string) {
	word, err := b.GetAt(x, y)
	if err != nil {
		return
	}
	if visit[x*b.Width()+y] {
		return
	}

	newWord := prefix + word
	if dict.IsWord(newWord) {
		ch <- newWord
	}
	if dict.IsPrefix(newWord) {
		visit[x*b.Width()+y] = true
		search(newWord, x-1, y, b, dict, visit, ch)
		search(newWord, x+1, y, b, dict, visit, ch)
		search(newWord, x, y-1, b, dict, visit, ch)
		search(newWord, x, y+1, b, dict, visit, ch)
		visit[x*b.Width()+y] = false
	}
}

func startSearch(x, y int, b *board.Board, dict *trie.Trie, ch chan string) {
	cap := b.Height() * b.Width()
	visit := make([]bool, cap, cap)
	// We'll ignore the error on the bootstrap
	word, err := b.GetAt(x, y)
	if err != nil {
		panic(err)
	}

	if dict.IsWord(word) {
		ch <- word
	}
	if dict.IsPrefix(word) {
		visit[x*b.Width()+y] = true
		search(word, x-1, y, b, dict, visit, ch)
		search(word, x+1, y, b, dict, visit, ch)
		search(word, x, y-1, b, dict, visit, ch)
		search(word, x, y+1, b, dict, visit, ch)
	}

	ch <- "ENDOFLINE"
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

	ch := make(chan string, 1000)

	for x := 0; x < b.Height(); x++ {
		for y := 0; y < b.Width(); y++ {
			go startSearch(x, y, b, dict, ch)
		}
	}

	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Printf("Solving took %v\n", elapsed)

	ends := 0
	var word string
	for {
		word = <-ch
		if "ENDOFLINE" == word {
			ends++
			if ends == 16 {
				break
			}
		} else {
			i, ok := words[word]
			if !ok {
				words[word] = 1
			} else {
				words[word] = i + 1
			}
		}
	}

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
