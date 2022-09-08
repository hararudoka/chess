package session

import (
	"errors"
)

// Ply is a half of a move. If white player moves his piece - that is the ply
type Ply struct {
	From Point
	To   Point
}

// convert Ply to string
func (p Ply) String() string {
	return p.From.String() + p.To.String()
}

// convert string to Ply
func (p *Ply) FromString(letters string) error {
	if len(letters) != 4 {
		return errors.New("invalid ply: " + letters)
	}
	from, to := Point{}, Point{}
	err := from.FromString(letters[:2])
	if err != nil {
		return err
	}
	err = to.FromString(letters[2:])
	if err != nil {
		return err
	}
	p.From = from
	p.To = to
	return nil
}

func MustFromString(letters string) Ply {
	p := Ply{}
	err := p.FromString(letters)
	if err != nil {
		panic(err)
	}
	return p
}
