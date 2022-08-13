package session

import (
	"errors"
	"fmt"
)

type Board [8][8]Piece

// move moves piece without checking if it's possible. it is the low-level method.
func (b Board) move(p Ply) Piece {
	b[p.To.File][p.To.Rank] = b[p.From.File][p.From.Rank]
	b[p.From.File][p.From.Rank] = Piece{}
	return b[p.To.File][p.To.Rank]
}

func (r Round) GetPossiblePoints(point Point, side Side, rank string) Points {
	if rank == Knight {
		coords, err := r.Knight(point, side)
		if err != nil {
			return []Point{}
		}
		return coords
	}
	// if rank == Pawn {
	// 	coords, err := r.Pawn(point)
	// 	if err != nil {
	// 		return []Point{}
	// 	}
	// 	return coords
	// }

	return Points{}
}

func (r Round) PieceCansGoHere(to Point, side Side, rank, x, y string) (Piece, error) {
	if x != "" && y != "" {
		point, err := StringToPoint(x + y)
		if err != nil {
			return Piece{}, err
		}
		return r.Board.GetPice(point), nil
	}

	if x == "" && y == "" {
		points := r.GetPossiblePoints(to, side, rank)

		fmt.Println("DEBUG:", points)

		for _, point := range points {
			fmt.Println("DEBUG:", r.Board.GetPice(point))
			if r.Board.GetPice(point).Rank() == rank {
				return r.Board.GetPice(point), nil
			}
		}
	}

	if x != "" {
		X := StringToY(x[0])
		for _, e := range r.Board[X] {
			if e.Colour() == side && e.Rank() == rank {
				return e, nil
			}
		}
	}
	if y != "" {
		Y := StringToX(x[0])
		for i, line := range r.Board {
			if File(i) == Y {
				for _, e := range line {
					if e.Colour() == side && e.Rank() == rank {
						return e, nil
					}
				}
			}
		}
	}
	return Piece{}, errors.New("not found piece")
}

func (b Board) GetPice(p Point) Piece {
	if p.File < 0 || p.File > 7 || p.Rank < 0 || p.Rank > 7 {
		return Piece{}
	}
	return b[p.File][p.Rank]
}
