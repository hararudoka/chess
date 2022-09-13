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
	Turn Turn
	// number of turn
	N int

	isLastX bool

	// castling
	CQ bool
	CK bool
	Cq bool
	Ck bool

	// last castling
	LastCastling string

	// enpassant
	EnPassant Point

	// halfmove clock to realisate 50 move rule
	HMW int
	HMB int

	// who is player
	SideOfPlayer Side
}

func (r *Round) Last() error {
	for {
		err := r.Next()
		if err != nil {
			return err
		}
	}
}


// play one turn
func (r *Round) Next() error {
	var err error

	if t := r.Turn.Get(r.Turn.Len() - r.N - 1); t != nil { // TODO: very tricky line, cuz len-n can be dangerous
		_, err = r.Ply(t.White, WhiteSide)
		if err != nil {
			return err
		}
		_, err = r.Ply(t.Black, BlackSide)
		if err != nil {
			return err
		}
		r.N++
	} else {
		return errors.New("no more turns")
	}
	return nil
}

func (r *Round) StartRecord() {
}

func (r *Round) FromFEN(fen string) error {
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
			return errors.New("invalid character in FEN: " + string(c))
		}
	}

	round := Round{}

	s := strings.Split(fen2, "/")

	if len(s) != 8 {
		return errors.New("wrong FEN string amout of rows")
	}

	ctt := strings.Split(cutten, " ")

	// who starts
	if ctt[0] == "w" {
		round.SideOfPlayer = WhiteSide
	} else if ctt[0] == "b" {
		round.SideOfPlayer = BlackSide
	} else {
		return errors.New("wrong FEN string")
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
		p := Point{}
		err := p.FromString(ctt[2])
		if err != nil {
			return err
		}
		round.EnPassant = p
	}

	// halfmove counter
	n, err := strconv.Atoi(ctt[3])
	if err != nil {
		return errors.New("wrong FEN string (halfmove counter white)")
	}
	round.HMW = n

	n, err = strconv.Atoi(ctt[4])
	if err != nil {
		return errors.New("wrong FEN string (halfmove counter black)")
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

	*r = round
	return nil
}

// makes ply
func (r *Round) Ply(ply Ply, side Side) (Piece, error) {
	ok, err := r.IsPlyPossible(ply, side)
	if err != nil {
		return Piece{}, err
	}
	if ok {
		p := r.Board.move(ply)
		return p, nil
	}
	return Piece{}, errors.New("ply is not possible " + ply.String())
}

// checks if ply is legal
func (r *Round) IsPlyPossible(ply Ply, side Side) (bool, error) {
	var err error
	var p Points

	side = side.Opposite()

	if ply.IsCastling {
		err = r.CastlingCheck(ply, side)
		if err != nil {
			return false, err
		}
		if side == WhiteSide {
			r.CK = false
			r.CQ = false
		} else {
			r.Ck = false
			r.Cq = false
		}
		return true, nil
	}

	switch r.Board.GetPice(ply.From).Class() {
	case "P":
		p, err = r.Pawn(ply.To, side)
	case "N":
		p, err = r.Knight(ply.To, side)
	case "B":
		p, err = r.Bishop(ply.To, side)
	case "R":
		p, err = r.Rook(ply.To, side)
	case "Q":
		p, err = r.Queen(ply.To, side)
	case "K":
		p, err = r.King(ply.To, side)
	default:
		return false, errors.New("wrong ply class")
	}
	if err != nil {
		return false, err
	}
	return p.contains(ply.From), nil
}

func (r Round) GetPossiblePoints(to Point, side Side, class string) Points {
	var points Points
	var err error

	switch class {
	case "P":
		points, err = r.Pawn(to, side)
	case "N":
		points, err = r.Knight(to, side)
	case "B":
		points, err = r.Bishop(to, side)
	case "R":
		points, err = r.Rook(to, side)
	case "Q":
		points, err = r.Queen(to, side)
	case "K":
		points, err = r.King(to, side)
	default:
		return Points{}
	}

	if err != nil {
		return Points{}
	}
	return points
}

// finds point from which piece can move to To point
func (r Round) FindFrom(to Point, side Side, class, rank, file string) (Point, error) {
	if rank != "" && file != "" {
		p := Point{}
		err := p.FromString(rank + file)
		if err != nil {
			return Point{}, err
		}
		return p, nil
	}

	if rank == "" && file == "" {
		points := r.GetPossiblePoints(to, side.Opposite(), class)
		// if class == "Q" {
		// 	panic(points)
		// }
		for _, point := range points {
			if r.Board.GetPice(point).Class() == class {
				return point, nil
			}
		}
	}

	if rank != "" {
		X := ByteToRank(rank[0])
		for i, e := range r.Board[X] {
			if e.Side() == side && e.Class() == class {
				return NewPoint(File(i), X), nil
			}
		}
	}

	if file != "" {
		if class == Pawn {
			points := r.GetPossiblePoints(to, side.Opposite(), class)
			for _, point := range points {
				if r.Board.GetPice(point).Class() == class {
					return point, nil
				}
			}
		}

		Y := ByteToFile(file[0])
		for i, line := range r.Board {
			if File(i) == Y {
				for x, e := range line {
					if e.Side() == side && e.Class() == class {
						return NewPoint(Y, Rank(x)), nil
					}
				}
			}
		}
	}

	return Point{}, errors.New("not found point for " + class + to.String())
}

func (r *Round) FromPGN(wholepgn, fen string) error {
	if fen == "" {
		fen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	}

	st := ""
	for _, e := range wholepgn {
		if e == '\n' { // TODO: add tabs and multiple spaces/\n
			st += " "
		} else {
			st += string(e)
		}
	}

	err := r.FromFEN(fen)
	if err != nil {
		return err
	}

	stringTurns := []string{}
	splitted := strings.Split(st, " ")
	for _, e := range splitted {
		// skip items that are not moves
		if e == "" {
			continue
		}
		if e == "." {
			continue
		}
		if _, err := strconv.Atoi(e); err == nil {
			continue
		}
		if _, err := strconv.Atoi(string(e[:len(e)-1])); err == nil && e[len(e)-1] == '.' {
			continue
		}

		if len(e) == 2 || len(e) == 3 || len(e) == 4 || len(e) == 5 || len(e) == 6 || len(e) == 7 {
			stringTurns = append(stringTurns, e)
		}
	}

	t, err := r.TurnFromStrings(stringTurns)
	if err != nil {
		return err
	}
	r.Turn = t
	return nil
}
