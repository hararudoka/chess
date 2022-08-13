package session

import (
	"errors"
	"strconv"
	"strings"
)

// Round is a state of board+some logix variables
type Round struct {
	// actual game board
	Board Board

	// list of moves
	Moves []Move

	// castling
	CQ bool
	CK bool
	Cq bool
	Ck bool

	// enpassant
	EnPassant Point

	// halfmove clock to realisate 50 move rule
	HMW int
	HMB int

	// who is player
	SideOfPlayer Side
}

func FENToRound(fen string) (Round, error) {
	fen2 := ""
	cutten := ""

loop:
	for i, c := range fen {
		switch c {
		case '8':
			fen2 += "11111111"
		case '7':
			fen2 += "1111111"
		case '6':
			fen2 += "111111"
		case '5':
			fen2 += "11111"
		case '4':
			fen2 += "1111"
		case '3':
			fen2 += "111"
		case '2':
			fen2 += "11"
		case '1':
			fen2 += "1"
		case '/':
			fen2 += "/"
		case 'p', 'n', 'b', 'r', 'q', 'k', 'P', 'N', 'B', 'R', 'Q', 'K':
			fen2 += string(c)
		case ' ':
			cutten = fen[i+1:]
			break loop
		default:
			return Round{}, errors.New("invalid character in FEN: " + string(c))
		}
	}

	round := Round{}

	s := strings.Split(fen2, "/")

	if len(s) != 8 {
		return Round{}, errors.New("wrong FEN string amout of rows")
	}

	ctt := strings.Split(cutten, " ")

	// who starts
	if ctt[0] == "w" {
		round.SideOfPlayer = White
	} else if ctt[0] == "b" {
		round.SideOfPlayer = Black
	} else {
		return Round{}, errors.New("wrong FEN string")
	}

	// castling rights
	for _, c := range ctt[1] {
		switch c {
		case 'K':
			round.CK = true
		case 'Q':
			round.CQ = true
		case 'k':
			round.Ck = true
		case 'q':
			round.Cq = true
		}
	}

	// en passant
	if ctt[2] != "-" {
		p, err := StringToPoint(ctt[2])
		if err != nil {
			return Round{}, err
		}
		round.EnPassant = p
	}

	// halfmove counter
	n, err := strconv.Atoi(ctt[3])
	if err != nil {
		return Round{}, errors.New("wrong FEN string (halfmove counter white)")
	}
	round.HMW = n

	n, err = strconv.Atoi(ctt[4])
	if err != nil {
		return Round{}, errors.New("wrong FEN string (halfmove counter black)")
	}
	round.HMB = n

	board := Board{}
	for i, row := range s {
		for j, c := range row {
			if c == '1' {
				continue
			}
			board[i][j] = getPiece(string(c))
		}
	}

	round.Board = board

	return round, nil
}

// RawPly processes input and moves piece to new position if cans
func (r Round) RawPly(raw string) error {
	p, _ := ToPGNPly(raw)
	ply, _ := r.PGNPlyToPly(p, r.SideOfPlayer)
	r.Board.move(ply)
	return nil
}
