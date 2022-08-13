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
func (p Point) ToLetters() string {
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

func StringToPoint(letters string) (Point, error) {
	if len(letters) != 2 {
		return Point{}, errors.New("letters must be 2 characters long")
	}

	// reversed
	x := StringToX(letters[1])
	y := StringToY(letters[0])

	return NewPoint(x, y), nil
}
