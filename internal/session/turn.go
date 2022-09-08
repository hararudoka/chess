package session

import "errors"

// Turn is one full cycle two players' board interaction
type Turn struct {
	next, prev *Turn

	White Ply
	Black Ply
}

func (turn *Turn) Add(newTurn Turn) bool {
	// if turn have data in next field it is impossible to add another turn (actually it is possible but i don't need it)
	if turn.next != nil {
		return false
	}
	// if newTurn has more than one turn it is impossible to add
	if newTurn.prev != nil {
		return false
	}
	// if turn is empty
	if turn.Black == (Ply{}) && turn.White == (Ply{}) {
		*turn = newTurn
		return true
	}
	// if prev is nil just move turn to prev and add data from newTurn
	if turn.prev == nil {
		turn.prev = turn.Copy()
		turn.White = newTurn.White
		turn.Black = newTurn.Black
		return true
	}

	cp := &Turn{}
	cp = turn.prev.Copy()

	*turn = newTurn
	turn.prev = cp

	return true
}

func (t Turn) Copy() *Turn {
	return &t
}

func (t *Turn) String() string {
	if t.prev == nil {
		return t.White.String() + "+" + t.Black.String()
	}
	return t.prev.String() + " -> " + t.White.String() + "+" + t.Black.String()
}

func (turn *Turn) Equal(newTurn Turn) bool {
	if turn.White != newTurn.White {
		return false
	}
	if turn.Black != newTurn.Black {
		return false
	}
	if turn.prev == nil && newTurn.prev == nil {
		return true
	}
	if turn.prev == nil || newTurn.prev == nil {
		return false
	}

	return turn.prev.Equal(*newTurn.prev)
}

// return root of the move list from strings
func (r Round) TurnFromStrings(turns []string) (Turn, error) {
	if len(turns)%2 == 1 {
		return Turn{}, errors.New("odd number of moves")
	}

	t := Turn{}
	for i := 0; i < len(turns); i += 2 {
		ply1, err := r.PlyFromString(turns[i], WhiteSide)
		if err != nil {
			return Turn{}, err
		}
		ply2, err := r.PlyFromString(turns[i+1], BlackSide)
		if err != nil {
			return Turn{}, err
		}

		r.Board.move(ply1)
		r.Board.move(ply2)

		t.Add(Turn{
			White: ply1,
			Black: ply2,
		})
	}

	return t, nil
}
