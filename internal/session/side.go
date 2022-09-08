package session

type Side string

func (side Side) Opposite() Side {
	if side == WhiteSide {
		return BlackSide
	}
	return WhiteSide
}

var (
	WhiteSide Side = "white"
	BlackSide Side = "black"
	EmptySide Side = ""
)
