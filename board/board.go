package board;

import (
  "fmt"
  "math/rand"
)

type BoardError struct {
	message string
}

func (e *BoardError) Error() string {
	return e.message
}

type Board struct {
  board [4][4]string;
  height int;
  width int;
}

// Gets a new random board
func New() *Board {
  dice := make([][]string, 16);
  dice[ 0] = []string{"a", "a", "c", "i", "o", "t"}; // AACIOT
  dice[ 1] = []string{"a", "h", "m", "o", "r", "s"}; // AHMORS
  dice[ 2] = []string{"e", "g", "k", "l", "u", "y"}; // EGKLUY
  dice[ 3] = []string{"a", "b", "i", "l", "t", "y"}; // ABILTY
  dice[ 4] = []string{"a", "c", "d", "e", "m", "p"}; // ACDEMP
  dice[ 5] = []string{"e", "g", "i", "n", "t", "v"}; // EGINTV
  dice[ 6] = []string{"g", "i", "l", "r", "u", "w"}; // GILRUW
  dice[ 7] = []string{"e", "l", "p", "s", "t", "u"}; // ELPSTU
  dice[ 8] = []string{"d", "e", "n", "o", "s", "w"}; // DENOSW
  dice[ 9] = []string{"a", "c", "e", "l", "r", "s"}; // ACELRS
  dice[10] = []string{"a", "b", "j", "m", "o", "q"}; // ABJMOQ
  dice[11] = []string{"e", "e", "f", "h", "i", "y"}; // EEFHIY
  dice[12] = []string{"e", "h", "i", "n", "p", "s"}; // EHINPS
  dice[13] = []string{"d", "k", "n", "o", "t", "u"}; // DKNOTU
  dice[14] = []string{"a", "d", "e", "n", "v", "z"}; // ADENVZ
  dice[15] = []string{"b", "i", "f", "o", "r", "x"}; // BIFORX

	var board Board;
  board.height = 4;
  board.width = 4;

  order := rand.Perm(16);
  pos := 0;
  for x := 0; x < board.height; x++ {
    for y := 0; y < board.width; y++ {
      board.board[x][y] = dice[order[pos]][rand.Int() % 6];
      pos++;
    }
  }

  return &board;
}

func (b *Board) Height() int {
  return b.height;
}

func (b *Board) Width() int {
  return b.width;
}

func (b *Board) GetAt(x, y int) (string, error) {
  if x < 0 || x >= b.height {
    return "", &BoardError{fmt.Sprintf("X value of %v is out of bounds must be 0 =< x < %v", x, b.height)};
  }
  if y < 0 || y >= b.width {
    return "", &BoardError{fmt.Sprintf("Y value of %v is out of bounds must be 0 =< y < %v", y, b.width)};
  }
  return b.board[x][y], nil;
}

func (b *Board) PrintBoard() {
  fmt.Println("Board:");
  for x := 0; x < b.height; x++ {
    for y := 0; y < b.width; y++ {
      fmt.Printf("%s ", b.board[x][y]);
    }
    fmt.Printf("\n");
  }
}
