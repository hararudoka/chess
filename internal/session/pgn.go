package session

import (
	"fmt"
	"strconv"
	"strings"
)

type PGNPly struct { // Longest example: Nb4xd5+!?
	// always first symbol or pawn
	Rank string

	// was it a capture
	IsTaked bool

	// can be 0, 1 or 2 symbols
	FileFrom string
	RankFrom string

	// 100% are 2 symbols
	To string

	IsCheck bool

	IsMate bool

	Descriptor string

	Special string
}

type Meta map[string]string

func (p PGNPly) String() string {
	return p.Rank + p.FileFrom + p.RankFrom + p.To + p.Special
}

func ToPGNPly(s string) (PGNPly, error) {
	ply := PGNPly{}

	if strings.Contains(s, "-") || s == "*" {
		ply.Special = s
		return ply, nil
	}

	if len(s) == 0 || len(s) == 1 {
		return ply, fmt.Errorf("invalid ply: %s", s)
	}

	if s[0] == 'N' || s[0] == 'B' || s[0] == 'R' || s[0] == 'Q' || s[0] == 'K' {
		ply.Rank = string(s[0])
	} else {
		ply.Rank = "P"
		s = "P" + s
	}

	ply.IsTaked = strings.Contains(s, "x")

	if ply.IsTaked {
		if strings.Index(s, "x") == 1 { // Nxd5 etc.
			ply.To = string(s[2]) + string(s[3])
		}
		if strings.Index(s, "x") == 2 { // Pexd5 Nexd5 etc.
			ply.To = string(s[3]) + string(s[4])

			if s[1] > 47 && s[1] < 58 { // digit
				ply.RankFrom = string(s[1])
			} else {
				ply.FileFrom = string(s[1])
			}
		}
		if strings.Index(s, "x") == 3 { // Pe3xd5 Ne1xd5 etc.
			ply.To = string(s[4]) + string(s[5])
			ply.FileFrom = string(s[1])
			ply.RankFrom = string(s[2])
		}
	}

	ply.IsCheck = strings.Contains(s, "+")
	ply.IsMate = strings.Contains(s, "#")

	ply.Descriptor = s[len(s)-1:]
	return ply, nil
}

func PGNToSession(pgn string) (Session, error) {
	meta, pgnCommentless := CommentsToList(pgn)

	round, err := PGNToRound(pgnCommentless, meta["FEN"])
	if err != nil {
		return Session{}, err
	}
	return Session{
		Round: round,
		Meta:  meta,
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

func PGNToRound(pgn, fen string) (Round, error) {
	if fen == "" {
		fen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	}

	st := ""
	for _, e := range pgn {
		if e == '\n' { // TODO: add tabs and multiple spaces/\n
			st += " "
		} else {
			st += string(e)
		}
	}
	r, err := FENToRound(fen)
	if err != nil {
		return r, err
	}
	moves := []string{}
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
			moves = append(moves, e)
		}
	}

	pgnplies := []PGNPly{}
	for i := 0; i < len(moves); i += 1 {
		p, err := ToPGNPly(moves[i])
		if err != nil {
			return Round{}, err
		}
		pgnplies = append(pgnplies, p)
	}

	r.PGNPliesToMoves(pgnplies)

	return r, nil
}

func (r *Round) PGNPliesToMoves(pgnplies []PGNPly) ([]Move, error) {
	for i := 0; i < len(pgnplies); i += 2 {
		wPly, err := r.PGNPlyToPly(pgnplies[i], White)
		if err != nil {
			return nil, err
		}
		bPly, err := r.PGNPlyToPly(pgnplies[i+1], Black)
		if err != nil {
			return nil, err
		}

		move := Move{
			White: wPly,
			Black: bPly,
		}

		r.Moves = append(r.Moves, move)
	}

	return r.Moves, nil
}

func (r Round) PGNPlyToPly(pgnPly PGNPly, side Side) (Ply, error) {
	p, err := StringToPoint(pgnPly.To)
	if err != nil {
		return Ply{}, err
	}

	r.PieceCansGoHere(p, side, pgnPly.Rank, pgnPly.RankFrom, pgnPly.FileFrom)

	return Ply{
		To: p,
	}, nil
}
