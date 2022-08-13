package session

// import (
// 	"testing"
// )

// // IMPORTANT:
// // colour of piece depends on the theme of your terminal
// // if your theme is white, then white pieces will be rendered as black, vice versa

// func TestCheck(t *testing.T) {
// 	tests := []struct {
// 		name  string
// 		start Session
// 		want  Session
// 		input []string
// 		err   error
// 	}{
// 		{
// 			name:  "knight pure moves",
// 			start: MustFENToSession("8/8/8/8/8/8/8/n7 w KQkq - 0 1"),
// 			want:  MustFENToSession("7n/8/8/8/8/8/8/8 w KQkq - 0 1"),
// 			input: []string{"a1c2", "c2e3", "e3c4", "c4e4", "e4g6", "g6h8"},
// 		},
// 		// {
// 		// 	name: "queen pure moves",
// 		// 	start: Session{
// 		// 		Board: FENToBoard(),
// 		// 	},
// 		// 	want: Session{
// 		// 		Board: FENToBoard(),
// 		// 	},
// 		// 	input: []Move{"A1H8", "H8A1", "A1A8", "A8H8", "H8H1", "H1A8", "A8A1", "A1H8", "H8H1"},
// 		// },
// 		// {
// 		// 	name: "queen wrong moves",
// 		// 	start: Session{
// 		// 		Board: FENToBoard(),
// 		// 	},
// 		// 	want: Session{
// 		// 		Board: FENToBoard(),
// 		// 	},
// 		// 	input: []Move{"A1A1"},
// 		// },
// 		// {
// 		// 	name: "pawn to queen moves",
// 		// 	start: Session{
// 		// 		Board: FENToBoard(),
// 		// 	},
// 		// 	want: Session{
// 		// 		Board: FENToBoard(),
// 		// 	},
// 		// 	input: []Move{"A2A4", "A4A5", "A5A6", "A6A7", "A7A8"},
// 		// },
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt := tt

// 			c := parseTestMoves(tt.input...)
// 			for _, ply := range c {
// 				_, log := tt.start.Move(ply)
// 				tt.want.Renderer.Add(log.Error())
// 			}

// 			if !isMapsEqual(tt.start.Board, tt.want.Board) {
// 				t.Errorf("Check()\nwant: \n%v\n got: \n%v", tt.want.Board, tt.start.Board)
// 			}
// 		},
// 		)
// 	}
// }

// func parseTestMoves(strings ...string) []Ply {
// 	var c []Ply
// 	for _, s := range strings {
// 		p, _ := StringToPly(s)
// 		c = append(c, p)
// 	}
// 	return c
// }

// // isMapsEqual returns true if the maps are equal
// func isMapsEqual(b1, b2 Board) bool {
// 	for i := 0; i < 8; i++ {
// 		for j := 0; j < 8; j++ {
// 			p1 := b1[i][j]
// 			p2 := b2[i][j]

// 			if p1 == (Piece{}) && p2 == (Piece{}) {
// 				continue
// 			}
// 			if p1 == (Piece{}) || p2 == (Piece{}) {
// 				return false
// 			}
// 			if p1.Kind != p2.Kind {
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }
