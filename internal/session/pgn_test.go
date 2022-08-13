package session

import (
	"fmt"
	"os"
	"testing"
)

func TestPGNToSession(t *testing.T) {
	tests := []struct {
		name    string
		pgnFile string
		want    Session
		err     error
	}{
		{
			name:    "test1",
			pgnFile: "../../test.pgn",
			want:    Session{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt

			// open file by name as string, copilot
			pgn, err := os.ReadFile(tt.pgnFile)
			if err != nil {
				panic(err)
			}
			got, err := PGNToSession(string(pgn))
			fmt.Println(got.Round.Moves[0])

			t.Errorf("PGNToSession()\nwant: \n%v\n got: \n%v", tt.want, got)
		},
		)
	}
}

func TestCommentsToList(t *testing.T) {
	tests := []struct {
		name    string
		pgnFile string
		want    map[string]string
		err     error
	}{
		{
			name:    "test1",
			pgnFile: "../../test.pgn",
			want: map[string]string{
				"Event": "abobus",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt

			// open file by name as string, copilot
			pgn, err := os.ReadFile(tt.pgnFile)
			if err != nil {
				panic(err)
			}

			got, _ := CommentsToList(string(pgn))
			if _, ok := got["Event"]; !ok {
				t.Errorf("CommentsToList()\nwant: \n%v\n got: \n%v", tt.want, got)
			}
		},
		)
	}
}

func TestPieceCansGoHere(t *testing.T) {
	point, _ := StringToPoint("g4")
	tests := []struct {
		name string

		fen string

		point      Point
		side       Side
		rank, x, y string

		want Piece
		err  error
	}{
		{
			name: "test1",

			fen: "8/8/8/8/8/5N2/8/8 w - - 0 1",

			point: point,
			side:  White,
			rank:  Knight,
			x:     "",
			y:     "",

			want: Piece{
				Kind:    "N",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt

			round, _ := FENToRound(tt.fen)
			got, err := round.PieceCansGoHere(tt.point, tt.side, tt.rank, tt.x, tt.y)
			if err != nil {
				panic(err)
			}

			if got != tt.want {
				t.Errorf("CommentsToList()\nwant: \n%v\n got: \n%v", tt.want, got)
			}
		},
		)
	}
}
