package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"time"

	"github.com/mikehelmick/bogglego/pkg/board"
	"github.com/mikehelmick/bogglego/pkg/trie"
)

type pos struct {
	x, y int
}

var offsets = []pos{{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}}

func (p pos) add(o pos) pos {
	return pos{p.x + o.x, p.y + o.y}
}

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

func search(prefix string, p pos, b *board.Board, dict *trie.Trie, visit []bool, ch chan string) {
	word, err := b.GetAt(p.x, p.y)
	if err != nil {
		return
	}
	if visit[p.x*b.Width()+p.y] {
		return
	}

	newWord := prefix + word
	if dict.IsWord(newWord) {
		ch <- newWord
	}
	if dict.IsPrefix(newWord) {
		visit[p.x*b.Width()+p.y] = true
		for _, off := range offsets {
			search(newWord, p.add(off), b, dict, visit, ch)
		}
		visit[p.x*b.Width()+p.y] = false
	}
}

func startSearch(p pos, b *board.Board, dict *trie.Trie, ch chan string) {
	cap := b.Height() * b.Width()
	visit := make([]bool, cap, cap)
	// We'll ignore the error on the bootstrap
	word, err := b.GetAt(p.x, p.y)
	if err != nil {
		panic(err)
	}

	if dict.IsWord(word) {
		ch <- word
	}
	if dict.IsPrefix(word) {
		visit[p.x*b.Width()+p.y] = true
		for _, off := range offsets {
			search(word, p.add(off), b, dict, visit, ch)
		}
	}

	ch <- "ENDOFLINE"
}

func main() {
	seed := flag.Int("seed", 0, "random seed")
	filename := flag.String("file", "", "name of the dictionary file to load")
	flag.Parse()
	if *seed == 0 {
		rand.Seed(time.Now().UTC().UnixNano())
	} else {
		rand.Seed(int64(*seed))
	}
	log.Printf("%v", *filename)
	if *filename == "" {
		panic("You must provide a filename to load.")
	}
	dict := trie.New()
	loadDictionary(*filename, dict)

	b := board.New()
	b.PrintBoard()

	fmt.Println("Solving")
	start := time.Now()

	words := make(map[string]int)

	ch := make(chan string, 1000)

	for x := 0; x < b.Height(); x++ {
		for y := 0; y < b.Width(); y++ {
			go startSearch(pos{x, y}, b, dict, ch)
		}
	}

	t := time.Now()
	elapsed := t.Sub(start)

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
			words[word]++
		}
	}

	var keys []string
	for k := range words {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("%s : %v\n", k, words[k])
	}
	fmt.Printf("Found %v words\n", len(keys))
	fmt.Printf("Solving took %v\n", elapsed)
}
