package session

type Session struct {
	Round Round

	Meta Meta

	Category string
}

func New(pgn, category string) (Session, error) {
	session, err := PGNToSession(pgn)
	if err != nil {
		return Session{}, err
	}
	session.Category = category
	return session, nil
}

func NewFromFEN(fen, category string) (Session, error) {
	round, err := FENToRound(fen)
	if err != nil {
		return Session{}, err
	}
	session := Session{
		Round:    round,
		Category: category,
	}
	session.Category = category
	return session, nil
}
