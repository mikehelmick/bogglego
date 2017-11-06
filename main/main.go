package main

import (
  "bufio"
	"fmt"
	"github.com/mikehelmick/bogglego/trie"
  "os"
)

// Simple error check
func check(e error) {
  if e != nil {
    panic(e);
  }
}

func loadDictionary(filename string, dict *trie.Trie) {
  fmt.Printf("Loading dictionry from %s\n", filename);
  dat, err := os.Open(filename);
  check(err);
  defer dat.Close();

  scanner := bufio.NewScanner(dat)
  scanner.Split(bufio.ScanLines)

  var wordCount int32 = 0;
  for scanner.Scan() {
    dict.AddWord(scanner.Text());
    wordCount++;
    if wordCount % 1000 == 0 {
      fmt.Print(".");
    }
  }
  fmt.Printf("Loaded %v words.\n", wordCount);
}

func main() {
  args := os.Args[1:]
  if len(args) != 1 {
    panic("You must provide a filename to load.");
  }
  dict := trie.New();
  loadDictionary(args[0], dict);
}
