package session

import "strings"

type Piece struct {
	Kind string
}

func (p Piece) String() string {
	return p.Class() + " " + string(p.Side())
}

func (p Piece) Class() string {
	if p.Kind == "" {
		return Empty
	}
	return strings.ToUpper(p.Kind)
}

func (p Piece) IsEmpty() bool {
	return p.Kind == ""
}

func (p Piece) Value() int {
	if p.Kind == "P" || p.Kind == "p" {
		return 1
	}
	if p.Kind == "R" || p.Kind == "r" {
		return 5
	}
	if p.Kind == "Q" || p.Kind == "q" {
		return 9
	}
	if p.Kind == "K" || p.Kind == "k" {
		return 1000
	}
	if p.Kind == "B" || p.Kind == "b" {
		return 3
	}
	if p.Kind == "N" || p.Kind == "n" {
		return 3
	}
	return 0
}

func (p Piece) Side() Side {
	if p.Kind == "" {
		return EmptySide
	}
	if p.Kind == "P" || p.Kind == "N" || p.Kind == "B" || p.Kind == "R" || p.Kind == "Q" || p.Kind == "K" {
		return WhiteSide
	}
	return BlackSide
}

type Pieces []Piece

func getPiece(kind string) Piece {
	return Piece{
		Kind: kind,
	}
}

// Clases
var (
	Empty  = ""
	King   = "K"
	Queen  = "Q"
	Rook   = "R"
	Bishop = "B"
	Knight = "N"
	Pawn   = "P"
)
