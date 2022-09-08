package render

import (
	"github.com/hararudoka/chess/internal/session"
)

func (r *Render) Chess() {
	// s, err := session.New("*", "chess")
	// if err != nil {
	// 	r.ErrorLine(err.Error())
	// 	return
	// }
	// s.Round.SideOfPlayer = session.WhiteSide
	// for {
	// 	clear()

	// 	board := r.RenderBoard(s.Round.Board)
	// 	r.Print(board)

	// 	input, err := r.Scan()
	// 	if err != nil {
	// 		r.ErrorLine(err.Error())
	// 		continue
	// 	}

	// 	s.Round.Ply(input)
	// }
}

func (r *Render) FEN() {
	round := session.Round{}
	for {
		clear()

		r.Print(r.RenderBoard(round.Board))

		fen, err := r.Scan()
		if err != nil {
			r.ErrorLine(err.Error())
			continue
		}

		err = round.FromFEN(fen)
		if err != nil {
			r.ErrorLine(err.Error())
			continue
		}
	}
}

func (r *Render) PGN() {
	clear()
	pgn := `1. e4 e5`

	s, err := session.New(pgn, "chess")
	if err != nil {
		r.ErrorLine(err.Error())
	}

	board := r.RenderBoard(s.Round.Board)
	r.Print(board)

	clear()

	s.StartRecord()

	r.Print(board)
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

		board := r.RenderBoard(s.Round.Board)
		r.Print(board)

		input, err := r.Scan()
		if err != nil {
			r.ErrorLine(err.Error())
			continue
		}

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
		r.Print(r.RenderBoard(round.Board))

		input, err := r.Scan()
		if err != nil {
			r.ErrorLine(err.Error())
			continue
		}
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
