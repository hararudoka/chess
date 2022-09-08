package session

import (
	"fmt"
	"os"
	"testing"
)

func TestNewFromPGN(t *testing.T) {
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
			got, err := NewFromPGN(string(pgn))
			fmt.Println(got.Round.Turn)

			t.Errorf("NewFromPGN()\nwant: \n%v\n got: \n%v", tt.want, got)
		},
		)
	}
}

func TestCommentsToList(t *testing.T) {
	tests := []struct {
		name    string
		pgnFile string
		want    Meta
		err     error
	}{
		{
			name:    "test1",
			pgnFile: "../../test.pgn",
			want: Meta{
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
			if _, ok := got["Event"]; !ok { // was i drunk when i wrote this?
				t.Errorf("CommentsToList()\nwant: \n%v\n got: \n%v", tt.want, got)
			}
		},
		)
	}
}

func TestPlyFromString(t *testing.T) {
	tests := []struct {
		name string

		fen string

		input string

		want Ply

		side Side

		err error
	}{
		{
			name: "night one move",

			fen: "8/8/8/8/8/5N2/8/8 w - - 0 1",

			side: WhiteSide,

			input: "Ng5",

			want: MustFromString("f3g5"),
		},
		{
			name: "queen white diagonal",

			fen: "8/8/8/8/8/8/8/7Q w - - 0 1",

			side: WhiteSide,

			input: "Qa8",

			want: MustFromString("h1a8"),
		},
		{
			name: "few rooks",

			fen: "8/8/8/8/8/8/8/RRRRRRRR w - - 0 1",

			side: WhiteSide,

			input: "Ra8",

			want: MustFromString("a1a8"),
		},
		{
			name: "few queens",

			fen: "8/8/8/8/8/8/8/Q1Q1Q1Q1 w - - 0 1",

			side: WhiteSide,

			input: "Qa8",

			want: MustFromString("a1a8"),
		},
		{
			name: "bishop at the center",

			fen: "8/8/8/4B3/8/8/8/8 w - - 0 1",

			side: WhiteSide,

			input: "Ba1",

			want: MustFromString("e5a1"),
		},
		{
			name: "pawns",

			fen: "8/8/8/8/8/8/PPPPPPPP/8 w - - 0 1",

			side: WhiteSide,

			input: "a3",

			want: MustFromString("a2a3"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt

			round := Round{}
			err := round.FromFEN(tt.fen)
			if err != nil {
				panic(err)
			}

			got, err := round.PlyFromString(tt.input, tt.side)
			if err != tt.err {
				t.Errorf("PlyFromString()\nwant: \n%v\n got: \n%v", tt.err, err)
			}

			point := Point{}
			point.FromString("h1")

			if got != tt.want {
				t.Errorf("PlyFromString()\nwant: \n%v\n got: \n%v", tt.want, got)
			}
		},
		)
	}
}

func TestTurnFromStrings(t *testing.T) {
	tests := []struct {
		name string

		input []string
		want  Turn
		err   error
	}{
		{
			name: "nights one move",
			input: []string{
				"Nf3",
				"Nf6",
			},
			want: Turn{
				White: MustFromString("g1f3"),
				Black: MustFromString("g8f6"),
			},
		},
		{
			name: "nights two moves 1",
			input: []string{
				"Nf3",
				"Nf6",
				"Nh4",
				"Nh5",
			},
			want: Turn{
				prev: &Turn{
					White: MustFromString("g1f3"),
					Black: MustFromString("g8f6"),
				},

				White: MustFromString("f3h4"),
				Black: MustFromString("f6h5"),
			},
		},
		{
			name: "nights two moves 2",
			input: []string{
				"Nf3",
				"Nf6",
				"Nd4",
				"Nd5",
			},
			want: Turn{
				prev: &Turn{
					White: MustFromString("g1f3"),
					Black: MustFromString("g8f6"),
				},

				White: MustFromString("f3d4"),
				Black: MustFromString("f6d5"),
			},
		},
		{
			name: "nights two moves 3",
			input: []string{
				"Nc3",
				"Na6",
				"Ne4",
				"Nb4",
			},
			want: Turn{
				prev: &Turn{
					White: MustFromString("b1c3"),
					Black: MustFromString("b8a6"),
				},

				White: MustFromString("c3e4"),
				Black: MustFromString("a6b4"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt

			round := Round{}
			err := round.FromFEN("rnbqkbn/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
			if err != nil {
				panic(err)
			}

			got, err := round.TurnFromStrings(tt.input)
			if err != nil {
				panic(err)
			}

			if !got.Equal(tt.want) {
				t.Errorf("TurnFromStrings()\nwant: \n%s\n got: \n%s", tt.want.String(), got.String())
			}
		},
		)
	}
}
