package session

// Board is a chess board. It is a 2d array of Pieces. Rank is Board[Rank], File is Board[n][File]
type Board [8][8]Piece

// move moves piece without checking if it's possible. it is the low-level method.
func (b *Board) move(p Ply) Piece {
	b[p.To.Rank][p.To.File] = b[p.From.Rank][p.From.File]
	b[p.From.Rank][p.From.File] = Piece{}
	return b[p.To.Rank][p.To.File]
}

// GetPice returns Piece on given Point. if Point is wrong returns empty Piece, so there is no guarantee of valid output
func (b Board) GetPice(p Point) Piece {
	if p.File < 0 || p.File > 7 || p.Rank < 0 || p.Rank > 7 {
		return Piece{}
	}
	return b[p.Rank][p.File]
}
