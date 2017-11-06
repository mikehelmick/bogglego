package main;

import(
  "fmt"
  "github.com/mikehelmick/bogglego/trie"
)

func main() {
  fmt.Println("Hello, world!")

  dictionary := trie.New()
  fmt.Println(dictionary);
}
