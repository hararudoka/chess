package session

import (
	"testing"
)

// IMPORTANT:
// colour of piece depends on the theme of your terminal
// if your theme is white, then white pieces will be rendered as black, vice versa

func TestLast(t *testing.T) {
	s := Session{}
	tests := []struct {
		name  string
		start Round
		want  Round
		input string
		err   error
	}{
		{
			name:  "pawn take",
			start: s.Round.FromFEN("8/8/8/8/8/8/8/n7 w KQkq - 0 1"),
			want:  MustFENToRound("7n/8/8/8/8/8/8/8 w KQkq - 0 1"),
			input: "1. e4 d5 2. exd5 Qd5",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt

			if !areBoardsEqual(tt.start.Board, tt.want.Board) {
				t.Errorf("Check()\nwant: \n%v\n got: \n%v", tt.want.Board, tt.start.Board)
			}
		},
		)
	}
}

func parseTestMoves(strings ...string) []Ply {
	var c []Ply
	for _, s := range strings {
		p, _ := StringToPly(s)
		c = append(c, p)
	}
	return c
}

// areBoardsEqual returns true if the maps are equal
func areBoardsEqual(b1, b2 Board) bool {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			p1 := b1[i][j]
			p2 := b2[i][j]

			if p1 == (Piece{}) && p2 == (Piece{}) {
				continue
			}
			if p1 == (Piece{}) || p2 == (Piece{}) {
				return false
			}
			if p1.Kind != p2.Kind {
				return false
			}
		}
	}
	return true
}