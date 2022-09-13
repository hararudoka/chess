package session

import (
	"fmt"
	"strings"
)

type Meta map[string]string

type pgn struct { // Longest example: Nb4xd5+!?
	// always first symbol or pawn
	Class string // N

	// was it a capture
	IsTaked bool // x

	// can be 0, 1 or 2 symbols
	FileFrom string // a
	RankFrom string // 1

	// 100% are 2 symbols
	To string // d5

	IsCheck bool // +

	IsMate bool // #

	Descriptor string // ?

	Special string // O-O-O
}

func (r *Round) fillpgn(s string) (pgn, error) {
	p := pgn{}
	// check on special
	if strings.Contains(s, "-") || s == "*" {
		p.Special = s
		return p, nil
	}

	s = strings.TrimSpace(s)

	// check on errors
	if len(s) == 0 || len(s) == 1 {
		return pgn{}, fmt.Errorf("invalid lenght: %s", s)
	}

	p.IsTaked = strings.Contains(s, "x")
	p.IsCheck = strings.Contains(s, "+")
	p.IsMate = strings.Contains(s, "#")

	if s[0] == 'N' || s[0] == 'B' || s[0] == 'R' || s[0] == 'Q' || s[0] == 'K' {
		p.Class = string(s[0])
	} else {
		p.Class = "P"
		s = "P" + s
	}

	if len(s) == 2 {
		return pgn{}, fmt.Errorf("invalid lenght: %s", s)
	}

	if len(s) == 3 {
		p.To = string(s[1]) + string(s[2])
		return p, nil
	}

	if p.IsTaked {
		r.isLastX = true
		if strings.Index(s, "x") == 1 { // Nxd5 etc.
			p.To = string(s[2]) + string(s[3])
		}
		if strings.Index(s, "x") == 2 { // Pexd5 Nexd5 etc.
			p.To = string(s[3]) + string(s[4])

			if s[1] > 47 && s[1] < 58 { // digit
				p.RankFrom = string(s[1])
			} else {
				p.FileFrom = string(s[1])
			}

		}
		if strings.Index(s, "x") == 3 { // Pe3xd5 Ne1xd5 etc.
			p.To = string(s[4]) + string(s[5])
			p.FileFrom = string(s[1])
			p.RankFrom = string(s[2])
		}
	} else {
		r.isLastX = false
	}

	// if s == "Qxd5" {
	// 	panic(p.To)
	// }

	return p, nil
}

// converts a ply written in PGN to a Ply
func (r *Round) PlyFromString(s string, side Side) (Ply, error) {
	pgn, err := r.fillpgn(s)
	if err != nil {
		return Ply{}, err
	}

	to := Point{}

	if pgn.Special != "" {
		if pgn.Special == "O-O" {
			if side == WhiteSide {
				r.LastCastling = "K"
				return Ply{
					IsCastling: true,
					From:       Point{4, 0},
					To:         Point{6, 0},
				}, nil
			}
			r.LastCastling = "k"
			return Ply{
				IsCastling: true,
				From:       Point{4, 7},
				To:         Point{6, 7},
			}, nil
		}
		if pgn.Special == "O-O-O" {
			if side == WhiteSide {
				r.LastCastling = "Q"
				return Ply{
					IsCastling: true,
					From:       Point{4, 0},
					To:         Point{2, 0},
				}, nil
			}
			r.LastCastling = "q"
			return Ply{
				IsCastling: true,
				From:       Point{4, 7},
				To:         Point{2, 7},
			}, nil
		}
	} else {
		if pgn.To == "" {
			panic("|" + s + "|")
		}
	}

	err = to.FromString(pgn.To)
	if err != nil {
		return Ply{}, err
	}

	from, err := r.FindFrom(to, side, pgn.Class, pgn.RankFrom, pgn.FileFrom)
	if err != nil {
		return Ply{}, err
	}

	return Ply{
		From: from,
		To:   to,
	}, nil
}

func CommentsToList(pgn string) (Meta, string) { // [Event "abobus"]
	insideBranch := false
	insideQuote := false

	moves := make(map[string]string)

	key, value := "", ""

	p := ""

	for _, e := range pgn {
		if e == '[' {
			insideBranch = true
			continue
		}

		if e == ']' {
			insideBranch = false
			continue
		}

		if e == ' ' && !insideBranch {
			p += string(e)
			continue
		}

		if e == ' ' && insideBranch && insideQuote {
			value += string(e)
			continue
		}

		if e == ' ' && insideBranch && !insideQuote {
			key += string(e)
			continue
		}

		if e == '"' && !insideQuote {
			insideQuote = true
			continue
		}
		if e == '"' && insideQuote {
			insideQuote = false
			continue
		}

		if insideBranch && !insideQuote {
			key += string(e)
			continue
		}

		if !insideBranch {
			p += string(e)
		}

		if insideBranch && insideQuote {
			value += string(e)
			continue
		}

		if key != "" {
			key = strings.TrimSpace(key)
			moves[key] = strings.TrimSpace(value)
		}
		key, value = "", ""

		if !insideBranch {
			continue
		}

	}

	return moves, p
}
