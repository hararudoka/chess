package session

import (
	"errors"
	"fmt"
)

// represents a tile on a board
type Point struct {
	File
	Rank
}

func NewPoint(x File, y Rank) Point {
	return Point{File(x), Rank(y)}
}

// changes map coordinates to "A1" like string
func (p Point) String() string {
	m := map[File]string{
		0: "A",
		1: "B",
		2: "C",
		3: "D",
		4: "E",
		5: "F",
		6: "G",
		7: "H",
	}
	return m[p.File] + fmt.Sprint(8-p.Rank)
}

func (p *Point) FromString(letters string) error {
	if len(letters) != 2 {
		return errors.New("letters must be 2 characters long: " + letters)
	}

	// reversed
	x := ByteToFile(letters[0])
	y := ByteToRank(letters[1])
	if x == -1 || y == -1 {
		return errors.New("invalid letters")
	}

	p.File = x
	p.Rank = y

	return nil
}
