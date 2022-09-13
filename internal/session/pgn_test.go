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

			input: []string{"Nf3", "Nf6"},
			want: Turn{
				White: MustFromString("g1f3"),
				Black: MustFromString("g8f6"),
			},
		},
		{
			name: "nights two moves 1",

			input: []string{"Nf3", "Nf6", "Nh4", "Nh5"},
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

			input: []string{"Nf3", "Nf6", "Nd4", "Nd5"},
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

			input: []string{"Nc3", "Na6", "Ne4", "Nb4"},
			want: Turn{
				prev: &Turn{
					White: MustFromString("b1c3"),
					Black: MustFromString("b8a6"),
				},

				White: MustFromString("c3e4"),
				Black: MustFromString("a6b4"),
			},
		},
		{
			name: "pawns d4 d5",

			input: []string{"d4", "d5"},
			want: Turn{
				White: MustFromString("d2d4"),
				Black: MustFromString("d7d5"),
			},
		},
		{
			name: "pawns d3 d6",

			input: []string{"d3", "d6"},
			want: Turn{
				White: MustFromString("d2d3"),
				Black: MustFromString("d7d6"),
			},
		},
		{
			name: "pawns a3 a6",

			input: []string{"a3", "h6", "a4", "h5", "a5", "h4", "a6", "h3"},
			want: Turn{
				prev: &Turn{
					prev: &Turn{
						prev: &Turn{
							White: MustFromString("a2a3"),
							Black: MustFromString("h7h6"),
						},
						White: MustFromString("a3a4"),
						Black: MustFromString("h6h5"),
					},
					White: MustFromString("a4a5"),
					Black: MustFromString("h5h4"),
				},
				White: MustFromString("a5a6"),
				Black: MustFromString("h4h3"),
			},
		},
		{
			name: "bishops",

			input: []string{"d4", "d5", "Bf4", "Bf5", "Nf3", "Nf6"},
			want: Turn{
				prev: &Turn{
					prev: &Turn{
						White: MustFromString("d2d4"),
						Black: MustFromString("d7d5"),
					},
					White: MustFromString("c1f4"),
					Black: MustFromString("c8f5"),
				},
				White: MustFromString("g1f3"),
				Black: MustFromString("g8f6"),
			},
		},
		{
			name: "pawn take",

			input: []string{"e4", "d5", "exd5", "Nf6"},
			want: Turn{
				prev: &Turn{
					White: MustFromString("e2e4"),
					Black: MustFromString("d7d5"),
				},
				White: MustFromString("e4d5"),
				Black: MustFromString("g8f6"),
			},
		},
		{
			name: "pawn take queen take",

			input: []string{"e4", "d5", "exd5", "Qxd5"},
			want: Turn{
				prev: &Turn{
					White: MustFromString("e2e4"),
					Black: MustFromString("d7d5"),
				},
				White: MustFromString("e4d5"),
				Black: MustFromString("d8d5"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			var gottenErr error

			round := Round{}
			err := round.FromFEN("rnbqkbn/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
			if err != nil {
				panic(err)
			}

			got, err := round.TurnFromStrings(tt.input)
			if err != nil {
				gottenErr = err
			}

			if tt.name == "pawn take" {
				r := round
				r.Last()
				fmt.Println(r.Board.GetPice(Point{3, 1}))
			}

			if gottenErr != tt.err {
				t.Errorf("err := TurnFromStrings()\nwant: %v\n got: %v", tt.err, gottenErr)
			} else if !got.Equal(tt.want) {
				t.Errorf("turn := TurnFromStrings()\nwant: %s\n got: %s", tt.want.String(), got.String())
			}
		},
		)
	}
}
