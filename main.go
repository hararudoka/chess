package main

import "fmt"

type Pices []Pice

type Pice struct {
	ID int // ID of piece

	// types:
	// 1 - pawn
	// 2 - knight
	// 3 - bishop
	// 4 - rook
	// 5 - queen
	// 6 - king
	Type int

	// sides: true - white, false - black
	Side bool // side of piece

	Value int
}

// create every piece
func GeneratePices() Pices {
	pices := Pices{Pice{ID: 0}}
	for i := 0; i < 16; i++ {
		if i < 8 {
			pices = append(pices, Pice{ID: i + 1, Type: 1, Side: true, Value: 1})
		} else {
			pices = append(pices, Pice{
				ID:   i + 1,
				Type: 1,
				Side: false,
			})
		}
	}

	grandPices := Pices{
		Pice{ID: 17, Type: 4},
		Pice{ID: 18, Type: 2},
		Pice{ID: 19, Type: 3},
		Pice{ID: 20, Type: 5},
		Pice{ID: 21, Type: 6},
		Pice{ID: 22, Type: 3},
		Pice{ID: 23, Type: 2},
		Pice{ID: 24, Type: 4},
		Pice{ID: 25, Type: 4},
		Pice{ID: 26, Type: 2},
		Pice{ID: 27, Type: 3},
		Pice{ID: 28, Type: 5},
		Pice{ID: 29, Type: 6},
		Pice{ID: 30, Type: 3},
		Pice{ID: 31, Type: 2},
		Pice{ID: 32, Type: 4},
	}

	pices = append(pices, grandPices...)
	return pices
}

func (ps Pices) GetTypeByID(id int) int {
	if id == 0 {
		return 0
	}
	return ps[id].Type
}

type Desk struct {
	Map [8][8]int
	Pices
}

func New() Desk {
	return Desk{
		Map: [8][8]int{
			{17, 18, 19, 20, 21, 22, 23, 24},
			{1, 2, 3, 4, 5, 6, 7, 8},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{9, 10, 11, 12, 13, 14, 15, 16},
			{25, 26, 27, 28, 29, 30, 31, 32},
		},
		Pices: GeneratePices(),
	}
}

func (d Desk) GetTypeByCords(x, y int) int {
	return d.Pices[d.Map[x][y]].Type
}

// return 1 == taked
// return 0 == not taked
// return -1 == move is not possible
func (d Desk) MoveByCords(x1, y1, x2, y2 int) int {
	if d.Map[x2][y2] != 0 { // TODO: after creation of players need to add check for side and taking of piece
		return -1
	}

	d.Map[x2][y2] = d.Map[x1][x1]
	d.Map[x1][x1] = 0

	return 1
}

func main() {
	d := New()

	fmt.Print(d.MoveByCords(0, 0, 3, 3))
}
