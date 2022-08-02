package desk

import "fmt"

type coords [][2]int

func (c coords) contains(x, y int) bool {
	var iterations int
	for left, right := 0, len(c)-1; left < len(c)/2; {
		iterations++
		if c[left][0] == x && c[left][1] == y || c[right][0] == x && c[right][1] == y {
			return true
		}
		left++
		right--
	}

	if len(c)%2 != 0 {
		i := len(c) / 2
		if c[i][0] == x && c[i][1] == y {
			return true
		}
	}
	return false
}

// not sure about pointers
func (cords coords) removeWrong() {
	for i := 0; i < len(cords); i++ {
		if cords[i][0] < 0 || cords[i][0] > 7 || cords[i][1] < 0 || cords[i][1] > 7 { // out of range
			cords = append(cords[:i], cords[i+1:]...)
			i--
		}
	}
}

// knight possible moves
func knight(x, y int) coords {
	probs := coords{
		{x + 1, y + 2},
		{x + 1, y - 2},
		{x - 1, y + 2},
		{x - 1, y - 2},
		{x + 2, y + 1},
		{x + 2, y - 1},
		{x - 2, y + 1},
		{x - 2, y - 1},
	}
	probs.removeWrong()
	return probs
}

// pawn possible moves
func pawn(x, y int) coords { // TODO: add sides logic
	probs := coords{
		{x + 1, y},
		{x - 1, y},
		{x + 1, y + 1},
		{x - 1, y + 1},
	}
	probs.removeWrong()
	return probs
}

// bishop possible moves
func bishop(x, y int) coords {
	probs := coords{}

	for {
		x++
		y++
		if x > 7 || y > 7 || x < 0 || y < 0 {
			break
		}
		probs = append(probs, [2]int{x, y})
	}
	for {
		x--
		y--
		if x > 7 || y > 7 || x < 0 || y < 0 {
			break
		}
		probs = append(probs, [2]int{x, y})
	}
	for {
		x++
		y--
		if x > 7 || y > 7 || x < 0 || y < 0 {
			break
		}
		probs = append(probs, [2]int{x, y})
	}
	for {
		x--
		y++
		if x > 7 || y > 7 || x < 0 || y < 0 {
			break
		}
		probs = append(probs, [2]int{x, y})
	}
	probs.removeWrong()
	return probs
}

// rook possible moves
func rook(x, y int) coords {
	probs := coords{}

	for {
		x++
		if x > 7 || x < 0 {
			break
		}
		probs = append(probs, [2]int{x, y})
	}
	for {
		x--
		if x > 7 || x < 0 {
			break
		}
		probs = append(probs, [2]int{x, y})
	}
	for {
		y++
		if y > 7 || y < 0 {
			break
		}
		probs = append(probs, [2]int{x, y})
	}
	for {
		y--
		if y > 7 || y < 0 {
			break
		}
		probs = append(probs, [2]int{x, y})
	}
	probs.removeWrong()
	return probs
}

func queen(x, y int) coords {
	probs := rook(x, y)
	probs = append(probs, bishop(x, y)...)
	return probs
}

func king(x, y int) coords {
	probs := coords{
		{x + 1, y},
		{x - 1, y},
		{x, y + 1},
		{x, y - 1},
		{x + 1, y + 1},
		{x - 1, y + 1},
		{x + 1, y - 1},
		{x - 1, y - 1},
	}
	probs.removeWrong()
	return probs
}

// bruh
func generatePices() [8][8]*Pice {
	pB := &Pice{
		Display: "♙",
		Type:    "pawn",
		Colour:  "black",
	}
	pW := &Pice{
		Display: "♟",
		Type:    "pawn",
		Colour:  "white",
	}
	rB := &Pice{
		Display: "♖",
		Type:    "rook",
		Colour:  "black",
	}
	rW := &Pice{
		Display: "♜",
		Type:    "rook",
		Colour:  "white",
	}
	nB := &Pice{
		Display: "♘",
		Type:    "knight",
		Colour:  "black",
	}
	nW := &Pice{
		Display: "♞",
		Type:    "knight",
		Colour:  "white",
	}
	bB := &Pice{
		Display: "♗",
		Type:    "bishop",
		Colour:  "black",
	}
	bW := &Pice{
		Display: "♝",
		Type:    "bishop",
		Colour:  "white",
	}
	qB := &Pice{
		Display: "♕",
		Type:    "queen",
		Colour:  "black",
	}
	qW := &Pice{
		Display: "♛",
		Type:    "queen",
		Colour:  "white",
	}
	kB := &Pice{
		Display: "♔",
		Type:    "king",
		Colour:  "black",
	}
	kW := &Pice{
		Display: "♚",
		Type:    "king",
		Colour:  "white",
	}

	return [8][8]*Pice{
		{rB, nB, bB, qB, kB, bB, nB, rB},         // 8
		{pB, pB, pB, pB, pB, pB, pB, pB},         // 7
		{nil, nil, nil, nil, nil, nil, nil, nil}, // 6
		{nil, nil, nil, nil, nil, nil, nil, nil}, // 5
		{nil, nil, nil, nil, nil, nil, nil, nil}, // 4
		{nil, nil, nil, nil, nil, nil, nil, nil}, // 3
		{pW, pW, pW, pW, pW, pW, pW, pW},         // 2
		{rW, nW, bW, qW, kW, bW, nW, rW},         // 1
		// A   B  C   D   E   F   G   H
	}
}

func lettersToDigits(x, y string) (int, int) {
	m := map[string]int{
		"A": 0,
		"B": 1,
		"C": 2,
		"D": 3,
		"E": 4,
		"F": 5,
		"G": 6,
		"H": 7,

		"1": 7,
		"2": 6,
		"3": 5,
		"4": 4,
		"5": 3,
		"6": 2,
		"7": 1,
		"8": 0,
	}
	return m[y], m[x]
}

func digitsToLetters(y, x int) string {
	m := map[int]string{
		0: "A",
		1: "B",
		2: "C",
		3: "D",
		4: "E",
		5: "F",
		6: "G",
		7: "H",
	}
	return m[x] + fmt.Sprint(8-y)
}
