package session

import (
	"errors"
)

// list of (x, y) pairs
type Points []Point

func (c Points) String() string {
	s := ""
	for _, e := range c {
		s += e.String() + ", "
	}
	return s
}

// contains checks if (x, y) is in the coords
func (ps Points) contains(p Point) bool {
	var iterations int
	for left, right := 0, len(ps)-1; left < len(ps)/2; {
		iterations++
		if ps[left] == p || ps[right] == p {
			return true
		}
		left++
		right--
	}

	if len(ps)%2 != 0 {
		i := len(ps) / 2
		if ps[i].File == p.File && ps[i].Rank == p.Rank {
			return true
		}
	}
	return false
}

// Tidy cleans Points from wrong ones
func (c *Points) Tidy() {
	c2 := make(Points, 0)
	for i, e := range *c {
		if c2.contains(e) {
			continue
		}
		if e.File <= 7 && e.File >= 0 && e.Rank <= 7 && e.Rank >= 0 {
			c2 = append(c2, (*c)[i])
		}
	}
	*c = c2
}

func (r *Round) Pawn(to Point, side Side) (Points, error) {
	var err error

	side = side.Opposite()

	points := Points{}

	if r.isLastX {
		if side == WhiteSide {
			if r.Board.GetPice(NewPoint(to.File+1, to.Rank-1)).Side() == side {
				points = append(points, NewPoint(to.File+1, to.Rank-1))
			}

			if r.Board.GetPice(NewPoint(to.File-1, to.Rank-1)).Side() == side {
				points = append(points, NewPoint(to.File-1, to.Rank-1))
			}
		}
		if side == BlackSide {
			if r.Board.GetPice(NewPoint(to.File+1, to.Rank+1)).Side() == side {
				points = append(points, NewPoint(to.File+1, to.Rank+1))
			}

			if r.Board.GetPice(NewPoint(to.File-1, to.Rank+1)).Side() == side {
				points = append(points, NewPoint(to.File-1, to.Rank+1))
			}
		}
	} else {
		if side == WhiteSide {
			points = append(points, NewPoint(to.File, to.Rank+1))

			p := Point{}
			p.FromString("a2")

			if to.Rank == p.Rank-2 {
				points = append(points, NewPoint(to.File, to.Rank+2))
			}
		}
		if side == BlackSide {
			points = append(points, NewPoint(to.File, to.Rank-1))

			p := Point{}
			p.FromString("a7")

			if to.Rank == p.Rank+2 {
				points = append(points, NewPoint(to.File, to.Rank-2))
			}
		}
	}

	points.Tidy()

	return points, err
}

// get allowed coords for each piece down below

func (r *Round) Knight(point Point, side Side) (Points, error) {
	var out Points
	var err error

	// dumply predicted
	points := Points{
		{point.File + 1, point.Rank + 2},
		{point.File + 1, point.Rank - 2},
		{point.File - 1, point.Rank + 2},
		{point.File - 1, point.Rank - 2},
		{point.File + 2, point.Rank + 1},
		{point.File + 2, point.Rank - 1},
		{point.File - 2, point.Rank + 1},
		{point.File - 2, point.Rank - 1},
	}

	// physical (not allowed cuz size of board) clean up
	points.Tidy()

	// logic clean up
	for _, e := range points {
		if r.Board.GetPice(e).Side() == side {
		} else {
			out = append(out, e)
		}
	}

	return out, err
}

func (r *Round) Bishop(point Point, side Side) (Points, error) {
	var out Points
	var err error

	// + +
	for i := 0; i < 8; i++ {
		// physical (not allowed cuz size of board) clean up
		if point.File+File(i) > 7 || point.Rank+Rank(i) > 7 {
			break
		}

		p := NewPoint(point.File+File(i), point.Rank+Rank(i))

		// logic clean up
		if r.Board.GetPice(p).Side() == side {
			break
		} else if r.Board.GetPice(p).Side() == EmptySide {
			out = append(out, p)
		} else {
			out = append(out, p)
			break
		}
	}

	// - -
	for i := 0; i < 8; i++ {
		// physical (not allowed cuz size of board) clean up
		if point.File-File(i) < 0 || point.Rank-Rank(i) < 0 {
			break
		}
		p := NewPoint(point.File-File(i), point.Rank-Rank(i))

		// logic clean up
		if r.Board.GetPice(p).Side() == side {
			break
		} else if r.Board.GetPice(p).Side() == EmptySide {
			out = append(out, p)
		} else {
			out = append(out, p)
			break
		}
	}

	// + -
	for i := 0; i < 8; i++ {
		// physical (not allowed cuz size of board) clean up
		if point.File+File(i) > 7 || point.Rank-Rank(i) < 0 {
			break
		}
		p := NewPoint(point.File+File(i), point.Rank-Rank(i))

		// logic clean up
		if r.Board.GetPice(p).Side() == side {
			break
		} else if r.Board.GetPice(p).Side() == EmptySide {
			out = append(out, p)
		} else {
			out = append(out, p)
			break
		}
	}

	// - +
	for i := 0; i < 8; i++ {
		// physical (not allowed cuz size of board) clean up
		if point.File-File(i) < 0 || point.Rank+Rank(i) > 7 {
			break
		}
		p := NewPoint(point.File-File(i), point.Rank+Rank(i))

		// logic clean up
		if r.Board.GetPice(p).Side() == side {
			break
		} else if r.Board.GetPice(p).Side() == EmptySide {
			out = append(out, p)
		} else {
			out = append(out, p)
			break
		}
	}

	out.Tidy()

	return out, err
}

func (r *Round) Rook(to Point, side Side) (Points, error) {
	var out Points
	var err error

	// file +
	for i := 0; i < 8; i++ {
		// physical (not allowed cuz size of board) clean up
		if to.File+File(i) > 7 {
			break
		}
		p := NewPoint(to.File+File(i), to.Rank)

		// logic clean up
		if r.Board.GetPice(p).Side() == side {
			break
		} else if r.Board.GetPice(p).Side() == EmptySide {
			out = append(out, p)
		} else {
			out = append(out, p)
			break
		}
	}
	// file -
	for i := 0; i < 8; i++ {
		// physical (not allowed cuz size of board) clean up
		if to.File-File(i) < 0 {
			break
		}
		p := NewPoint(to.File-File(i), to.Rank)

		// logic clean up
		if r.Board.GetPice(p).Side() == side {
			break
		} else if r.Board.GetPice(p).Side() == EmptySide {
			out = append(out, p)
		} else {
			out = append(out, p)
			break
		}
	}
	// rank +
	for i := 0; i < 8; i++ {
		// physical (not allowed cuz size of board) clean up
		if to.Rank+Rank(i) > 7 {
			break
		}
		p := NewPoint(to.File, to.Rank+Rank(i))

		// logic clean up
		if r.Board.GetPice(p).Side() == side {
			break
		} else if r.Board.GetPice(p).Side() == EmptySide {
			out = append(out, p)
		} else {
			out = append(out, p)
			break
		}
	}
	// rank -
	for i := 0; i < 8; i++ {
		// physical (not allowed cuz size of board) clean up
		if to.Rank-Rank(i) < 0 {
			break
		}
		p := NewPoint(to.File, to.Rank-Rank(i))

		// logic clean up
		if r.Board.GetPice(p).Side() == side {
			break
		} else if r.Board.GetPice(p).Side() == EmptySide {
			out = append(out, p)
		} else {
			out = append(out, p)
			break
		}
	}
	out.Tidy()

	return out, err
}

// returns possible Points, which can be moved FROM
func (r *Round) Queen(to Point, side Side) (Points, error) {
	var out Points
	var err error

	bishop, err := r.Bishop(to, side)
	if err != nil {
		return out, err
	}

	rook, err := r.Rook(to, side)
	if err != nil {
		return out, err
	}

	out = bishop
	out = append(out, rook...)

	out.Tidy()

	//panic(r.Board.GetPice(to))

	return out, err
}

func (r *Round) King(point Point, side Side) (Points, error) {
	var out Points
	var err error

	// dumply predicted
	points := Points{
		{point.File + 1, point.Rank + 1},
		{point.File + 1, point.Rank - 1},
		{point.File - 1, point.Rank + 1},
		{point.File - 1, point.Rank - 1},
		{point.File, point.Rank + 1},
		{point.File, point.Rank - 1},
		{point.File + 1, point.Rank},
		{point.File - 1, point.Rank},
	}

	points.Tidy()

	for _, p := range points {
		if r.Board.GetPice(p).Side() == side {
		} else {
			out = append(out, p)
		}
	}

	return out, err
}

func (r *Round) CastlingCheck(point Ply, side Side) error {
	// TODO: checks
	c := r.LastCastling
	r.LastCastling = ""

	if c == "K" {
		if !r.CK {
			return errors.New("castling is dead")
		}

		// NOT WORKING YET
		if r.Board.GetPice(Point{6, 7}).Class() == Empty || r.Board.GetPice(Point{5, 7}).Class() == Empty {
			return errors.New("not allowed")
		}
	}
	if c == "Q" {
		if !r.CQ {
			return errors.New("castling is dead")
		}

		if r.Board.GetPice(Point{2, 7}).Class() == Empty || r.Board.GetPice(Point{3, 7}).Class() == Empty || r.Board.GetPice(Point{1, 7}).Class() == Empty {
			return errors.New("not allowed")
		}
	}
	if c == "k" {
		if !r.Ck {
			return errors.New("castling is dead")
		}

		if r.Board.GetPice(Point{6, 0}).Class() == Empty || r.Board.GetPice(Point{5, 0}).Class() == Empty {
			return errors.New("not allowed")
		}
	}
	if c == "q" {
		if !r.Cq {
			return errors.New("castling is dead")
		}

		if r.Board.GetPice(Point{2, 0}).Class() == Empty || r.Board.GetPice(Point{3, 0}).Class() == Empty || r.Board.GetPice(Point{1, 0}).Class() == Empty {
			return errors.New("not allowed")
		}
	}
	return nil
}
