package desk

import (
	"fmt"
)

type Pices []Pice

type Pice struct {
	Type    string
	Display string

	Colour string

	// Value int
}

//       The Desk
//    H G F E D C B A
// 1 |r|n|b|q|k|b|n|r| 1 // whites
// 2 |p|p|p|p|p|p|p|p| 2 // whites
// 3 | | | | | | | | | 3
// 4 | | | | | | | | | 4
// 5 | | | | | | | | | 5
// 6 | | | | | | | | | 6
// 7 |p|p|p|p|p|p|p|p| 7 // blacks
// 8 |r|n|b|q|k|b|n|r| 8 // blacks
//    H G F E D C B A

type Desk struct {
	Map   [8][8]*Pice
	Loses []*Pice
	Side  string
	Logs  Logs
}

func New() Desk {
	return Desk{
		// Usual desk
		Map:  generatePices(),
		Side: "white",
	}
}

// basic render of the desk
func (d Desk) String() string {
	// whites
	s := "  A B C D E F G H\n"
	for i, row := range d.Map {
		s += fmt.Sprint(8 - i)
		for _, pice := range row {
			if pice == nil {
				s += "| "
			} else {
				s += "|" + pice.Display
			}
		}
		s += "|" + fmt.Sprint(8-i) + "\n"
	}
	s += "  A B C D E F G H\n"
	return s
}

func (d *Desk) Check(x1, y1, x2, y2 int) bool {
	if x1 < 0 || x1 > 7 || y1 < 0 || y1 > 7 || x2 < 0 || x2 > 7 || y2 < 0 || y2 > 7 {
		d.Logs.Error(Log{String: "wrong coordinates"})
		return false
	}

	// logic depends on pice type
	if d.Map[x1][y1].Type == "knight" { // TODO: fix impossible moves
		if knight(x1, y1).contains(x2, y2) {
			return true
		}
		return false
	}

	if d.Map[x1][y1].Type == "pawn" { // TODO: rewrite logic
		if pawn(x1, y1).contains(x2, y2) {
			return true
		}
		return false
	}

	if d.Map[x1][y1].Type == "bishop" { // TODO: fix impossible moves
		if bishop(x1, y1).contains(x2, y2) {
			return true
		}
		return false
	}

	if d.Map[x1][y1].Type == "rook" { // TODO: fix impossible moves
		if rook(x1, y1).contains(x2, y2) {
			return true
		}
		return false
	}

	if d.Map[x1][y1].Type == "queen" { // TODO: fix impossible moves
		if queen(x1, y1).contains(x2, y2) {
			return true
		}
		return false
	}

	if d.Map[x1][y1].Type == "king" { // TODO: test king more carefully
		if king(x1, y1).contains(x2, y2) {
			return true
		}
		return false
	}

	if d.Map[x2][y2] == nil {
		return true
	}

	if d.Map[x2][y2].Colour == d.Map[x1][y1].Colour {
		return false
	}

	return true
}

func (d *Desk) Move(x1, y1, x2, y2 int) (bool, Log) { // TODO: errors
	if d.Map[x1][y1] == nil {
		fmt.Print("nil lol")
		return false, Log{String: "this is not piece"}
	}

	if d.Check(x1, y1, x2, y2) {
		d.Map[x2][y2] = d.Map[x1][y1]
		d.Map[x1][y1] = nil
		return true, Log{
			Side:   d.GetPice(x1, x2).Colour,
			From:   digitsToLetters(x1, y1),
			To:     digitsToLetters(x2, y2),
			String: fmt.Sprintf("%s: %s -> %s", d.GetPice(x2, y2).Colour, digitsToLetters(x1, y1), digitsToLetters(x2, y2)),
		}
	}
	return false, Log{String: "impossible move"}
}

func (d Desk) GetPice(x, y int) Pice {
	if d.Map[x][y] == nil {
		return Pice{Type: "empty"}
	}
	return *d.Map[x][y]
}

func (d Desk) Run(side string) {
	d.Side = side

	// theme
	fmt.Print("\x1b[48;5;0m\x1b[38;5;231m") // 0m == black, 231m == white

	// clear screen
	fmt.Print("\033[H\033[2J")
	d.Logs.Add(EmptyLog)

	for { // TODO: divide into functions and make it more readable
		// move coursor
		fmt.Print("\033[1;1H")
		fmt.Print(d)
		fmt.Print(">\n")
		fmt.Print(d.Logs.String())

		// move coursor behind board
		fmt.Print("\033[11;2H")

		// scan input
		move := ""
		fmt.Scan(&move)
		x1, y1 := lettersToDigits(string(move[0]), string(move[1]))
		x2, y2 := lettersToDigits(string(move[2]), string(move[3]))

		// clear screen
		fmt.Print("\033[H\033[2J")

		if ok, l := d.Move(x1, y1, x2, y2); ok {
			d.Logs.Add(l)
			d.Logs.List[0] = EmptyLog
		} else {
			d.Logs.Error(l)
		}

		// move coursor
		// fmt.Print("\033[11;1H")
	}
}
