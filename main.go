package main

import (
	"os"
)

type Board []string

func (b Board) solve() bool {
	for i, row := range b {
		for j, ch := range row {
			if ch != '.' {
				continue
			}

			var newCh uint8
			for newCh = '1'; newCh < '9'; newCh++ {
				b[i][j] = newCh
				if b.solve() {
					return true
				} else {
					b[i][j] = '.'
				}
			}
			return false
		}
	}
	return true
}

func (b Board) print() {
	for _, row := range b {
		for _, ch := range row {
			os.Stdout.Write([]byte{byte(ch)})
			os.Stdout.WriteString(" ")
		}
		os.Stdout.WriteString("\n")
	}
}

var board = []string{
"53..7....",
"6..195...",
".98....6.",
"8...6...3",
"4..8.3..1",
"7...2...6",
".6....28.",
"...419..5",
"....8..79",
}

func main() {
	b := Board(board)
	b.solve()
	b.print()
}
