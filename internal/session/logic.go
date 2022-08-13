package session

import "errors"

// list of (x, y) pairs
type Points []Point

func (c Points) String() string {
	s := ""
	for _, e := range c {
		s += e.ToLetters() + ", "
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

func (c *Points) Clean() {
	c2 := make(Points, 0)
	for i, e := range *c {
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
	switch piece.Kind {
	case Pawn:
		out, err = r.Pawn(point)
		// case Knight:
		// 	out, err = r.Knight(point)
		// case Bishop:
		// 	out, err = r.Bishop(point)
		// case Rook:
		// 	out, err = r.Rook(point)
		// case Queen:
		// 	out, err = r.Queen(point)
		// case King:
		// 	out, err = r.King(point)
	}

	if err != nil {
		return nil, err
	}
	possible = append(possible, out...)
	return possible, nil
}

func (r *Round) Pawn(point Point) (Points, error) {
	var out Points
	var err error

	return out, err
}

// get possible coords if knight starts from point
func (r *Round) Knight(point Point, side Side) (Points, error) {
	var out Points
	var err error

	out = Points{
		{point.File + 1, point.Rank + 2},
		{point.File + 1, point.Rank - 2},
		{point.File - 1, point.Rank + 2},
		{point.File - 1, point.Rank - 2},
		{point.File + 2, point.Rank + 1},
		{point.File + 2, point.Rank - 1},
		{point.File - 2, point.Rank + 1},
		{point.File - 2, point.Rank - 1},
	}
	out.Clean()

	// add more complex logic
	// sides, captures etc.

	return out, err
}
