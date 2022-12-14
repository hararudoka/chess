package session

type Session struct {
	Round Round

	Meta Meta

	Category string
}

func New(pgn, category string) (Session, error) {
	session, err := NewFromPGN(pgn)
	if err != nil {
		return Session{}, err
	}
	session.Category = category
	return session, nil
}

func NewFromPGN(pgn string) (Session, error) {
	meta, pgnCommentless := CommentsToList(pgn)

	round := Round{}
	err := round.FromPGN(pgnCommentless, meta["FEN"])
	if err != nil {
		return Session{}, err
	}
	return Session{
		Round: round,
		Meta:  meta,
	}, nil
}

func NewFromFEN(fen, category string) (Session, error) {
	round := Round{}
	err := round.FromFEN(fen)
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

func (s *Session) StartRecord() {
	s.Round.StartRecord()
}

func (s *Session) Move(input string) (Piece, error) {
	ply, err := s.Round.PlyFromString(input, s.Round.SideOfPlayer)
	if err != nil {
		return Piece{}, err
	}
	p, err := s.Round.Ply(ply, s.Round.SideOfPlayer)
	if err != nil {
		return Piece{}, err
	}
	return p, nil
}

func (s *Session) BotMove() (Piece, error) {
	// TODO: ...
	ply, err := s.Round.PlyFromString("a5", s.Round.SideOfPlayer.Opposite())
	if err != nil {
		return Piece{}, err
	}
	p, err := s.Round.Ply(ply, s.Round.SideOfPlayer.Opposite())
	if err != nil {
		return Piece{}, err
	}
	return p, nil
}
