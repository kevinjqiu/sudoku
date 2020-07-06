package main

import (
	"os"
)

const BoardSize = 9

type Board [BoardSize][BoardSize]uint8

func (b *Board) isValid(i, j int, newCh uint8) bool {
	for _, ch := range b[i] {
		if newCh == ch {
			return false
		}
	}

	for _, row := range b {
		if row[j] == newCh {
			return false
		}
	}

	for ii := 3 * (i / 3); ii < 3 * (i / 3) + 3; ii++ {
		for jj := 3 * (j / 3); jj < 3 * (j / 3) + 3; jj ++ {
			if b[ii][jj] == newCh {
				return false
			}
		}
	}
	return true
}

func (b *Board) solve() bool {
	for i, row := range b {
		for j, ch := range row {
			if ch != '.' {
				continue
			}

			var newCh uint8
			for newCh = '1'; newCh <= '9'; newCh++ {
				if b.isValid(i, j, newCh) {
					b[i][j] = newCh
					//b.print()
					if b.solve() {
						return true
					} else {
						b[i][j] = '.'
						//b.print()
					}
				}
			}
			return false
		}
	}
	return true
}

func (b Board) print() {
	for i, row := range b {
		for j, ch := range row {
			os.Stdout.Write([]byte{byte(ch)})
			os.Stdout.WriteString(" ")
			if (j+1) % 3 == 0 {
				os.Stdout.WriteString(" | ")
			}
		}
		if (i+1) % 3 == 0 {
			os.Stdout.WriteString("\n--------------------------\n")
		} else{
			os.Stdout.WriteString("\n")
		}
	}
	os.Stdout.WriteString("----------\n")
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

func makeBoard(rows []string) Board {
	b := Board{}

	for i, row := range rows {
		for j, ch := range row {
			b[i][j] = uint8(ch)
		}
	}

	return b
}

func main() {
	b := makeBoard(board)
	b.print()
	b.solve()
	b.print()
}
