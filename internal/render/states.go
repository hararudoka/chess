package render

import (
	"time"

	"github.com/hararudoka/chess/internal/session"
)

func (r *Render) Chess() {
	s, err := session.New("*", "chess")
	if err != nil {
		r.ErrorLine(err.Error())
		return
	}

	s.Round.SideOfPlayer = session.WhiteSide

	for {
		clear()

		board := r.StringBoard(s.Round.Board)
		r.Print(board)

		input := r.Scan()

		p, err := s.Move(input)
		if err != nil {
			r.ErrorLine(err.Error())
			continue
		}

		r.Add(input + " = " + p.String())
	}
}

func (r *Render) FEN() {
	round := session.Round{}
	for {
		clear()

		r.Print(r.StringBoard(round.Board))

		fen := r.Scan()

		err := round.FromFEN(fen)
		if err != nil {
			r.ErrorLine(err.Error())
			continue
		}
	}
}

func (r *Render) PGN() {
	clear()
	pgn := `1. e4 d5 2. exd5 Qd5`

	s, err := session.New(pgn, "chess")
	if err != nil {
		r.ErrorLine(err.Error())
		return
	}

	for {
		clear()

		r.PrintBoard(s.Round.Board)

		err := s.Round.Next()
		if err != nil {
			r.ErrorLine(err.Error())
		}

		// r.ErrorLine(fmt.Sprint(s.Round.Turn.Get(9)))

		time.Sleep(100 * time.Millisecond)
	}
}

func (r *Render) TestPly() {
	r.State = "test"
	clear()

	s, err := session.NewFromFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1", "test")
	if err != nil {
		r.ErrorLine(err.Error())
		return
	}

	for {
		clear()

		board := r.StringBoard(s.Round.Board)
		r.Print(board)

		input := r.Scan()

		ply, err := s.Round.PlyFromString(input, session.WhiteSide)
		if err != nil {
			r.ErrorLine(err.Error())
			continue
		}
		p, err := s.Round.Ply(ply, session.WhiteSide)
		if err != nil {
			r.ErrorLine(err.Error())
			continue
		}
		r.Add(input + " = " + p.String())
	}
}

func (r Render) TestGetPiece() {
	r.State = "test"
	clear()

	round := session.Round{}
	err := round.FromFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	if err != nil {
		panic(err)
	}

	for {
		clear()
		r.Print(r.StringBoard(round.Board))

		input := r.Scan()

		p := session.Point{}
		err = p.FromString(input)
		if err != nil {
			r.ErrorLine(err.Error())
			continue
		}

		s := round.Board.GetPice(p).Class() + " " + string(round.Board.GetPice(p).Side())
		if s == " " {
			r.Add(input + " = empty")
		} else {
			r.Add(input + " = " + s)
		}

	}
}
