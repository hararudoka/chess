package session

// Ply is a half of a move. If white player moves his piece - that is the ply
type Ply struct {
	From Point
	To   Point
}

// convert Ply to string
func (p Ply) ToLetters() string {
	return p.From.ToLetters() + p.To.ToLetters()
}

// convert string to Ply
func StringToPly(letters string) (Ply, error) {
	from, err := StringToPoint(letters[:2])
	if err != nil {
		return Ply{}, err
	}
	to, err := StringToPoint(letters[2:])
	if err != nil {
		return Ply{}, err
	}
	return Ply{
		From: from,
		To:   to,
	}, err
}
