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

func (r *Round) getPossibleCoords(point Point) (Points, error) {
	if point == (Point{}) { // TODO: what if point is not possible?
		return nil, errors.New("non existent point")
	}
	piece := r.Board.GetPice(point)
	if piece.IsEmpty() {
		return nil, errors.New("empty point")
	}

	var out Points
	var possible Points
	var err error
	switch piece.Class() {
	case Pawn:
		out, err = r.Pawn(point, piece.Side())
	case Knight:
		out, err = r.Knight(point, piece.Side())
	case Bishop:
		out, err = r.Bishop(point, piece.Side())
	case Rook:
		out, err = r.Rook(point, piece.Side())
	case Queen:
		out, err = r.Queen(point, piece.Side())
	case King:
		out, err = r.King(point, piece.Side())
	}

	if err != nil {
		return nil, err
	}
	possible = append(possible, out...)
	return possible, nil
}

func (r *Round) Pawn(point Point, side Side) (Points, error) {
	// var out Points
	var err error

	points := Points{
		NewPoint(point.File, point.Rank+1),
	}

	if side == WhiteSide {
		p := Point{}
		p.FromString("a2")
		if point.Rank == p.Rank {
			points = append(points, NewPoint(point.File, point.Rank+2))
		}
	}
	if side == BlackSide {
		p := Point{}
		p.FromString("a2")
		if point.Rank == p.Rank {
			points = append(points, NewPoint(point.File, point.Rank+2))
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

func (r *Round) Rook(point Point, side Side) (Points, error) {
	var out Points
	var err error

	// file +
	for i := 0; i < 8; i++ {
		// physical (not allowed cuz size of board) clean up
		if point.File+File(i) > 7 {
			break
		}
		p := NewPoint(point.File+File(i), point.Rank)

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
		if point.File-File(i) < 0 {
			break
		}
		p := NewPoint(point.File-File(i), point.Rank)

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
		if point.Rank+Rank(i) > 7 {
			break
		}
		p := NewPoint(point.File, point.Rank+Rank(i))

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
		if point.Rank-Rank(i) < 0 {
			break
		}
		p := NewPoint(point.File, point.Rank-Rank(i))

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

func (r *Round) Queen(point Point, side Side) (Points, error) {
	var out Points
	var err error

	bishop, err := r.Bishop(point, side)
	if err != nil {
		return out, err
	}

	rook, err := r.Rook(point, side)
	if err != nil {
		return out, err
	}

	out = bishop
	out = append(out, rook...)

	out.Tidy()

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
