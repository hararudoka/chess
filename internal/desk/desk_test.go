package desk

import (
	"testing"
)

func TestCheck(t *testing.T) {
	qW := &Pice{
		Type:   "queen",
		Colour: "white",
	}

	tests := []struct {
		name  string
		desk  Desk
		want  Desk
		input []int
		err   error
	}{
		{
			name: "queen",
			desk: Desk{
				Map: [8][8]*Pice{
					{qW, nil, nil, nil, nil, nil, nil, nil},
					{nil, nil, nil, nil, nil, nil, nil, nil},
					{nil, nil, nil, nil, nil, nil, nil, nil},
					{nil, nil, nil, nil, nil, nil, nil, nil},
					{nil, nil, nil, nil, nil, nil, nil, nil},
					{nil, nil, nil, nil, nil, nil, nil, nil},
					{nil, nil, nil, nil, nil, nil, nil, nil},
					{nil, nil, nil, nil, nil, nil, nil, nil},
				},
			},
			want: Desk{
				Map: [8][8]*Pice{
					{nil, nil, nil, nil, nil, nil, nil, nil},
					{nil, nil, nil, nil, nil, nil, nil, nil},
					{nil, nil, nil, nil, nil, nil, nil, nil},
					{nil, nil, nil, nil, nil, nil, nil, nil},
					{nil, nil, nil, nil, nil, nil, nil, nil},
					{nil, nil, nil, nil, nil, nil, nil, nil},
					{nil, nil, nil, nil, nil, nil, nil, nil},
					{qW, nil, nil, nil, nil, nil, nil, nil},
				},
			},
			input: []int{0, 0, 7, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt

			tt.desk.Move(tt.input[0], tt.input[1], tt.input[2], tt.input[3])
			// if err != tt.err {
			// 	t.Errorf("SanitizeURL()\nwant error: %v\n got error: %v", tt.err, err)
			// 	return
			// }
			got := tt.desk
			if got.Map != tt.want.Map {
				t.Errorf("Check()\nwant: %v\n got: %v", tt.want, got.Map)
			}
		},
		)
	}
}
