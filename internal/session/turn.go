package session

import (
	"errors"
)

// Turn is one full cycle two players' board interaction
type Turn struct {
	prev *Turn

	White Ply
	Black Ply
}

func (t *Turn) Len() int {
	if t == nil {
		return 0
	}
	return 1 + t.prev.Len()
}

func (t *Turn) Get(n int) *Turn {
	if t == nil {
		return nil
	}
	if n == 0 {
		return t
	}
	return t.prev.Get(n - 1)
}

func (t *Turn) First() *Turn {
	if t == nil {
		return nil
	}
	if t.prev == nil {
		return t
	}
	return t.prev.First()
}

func (turn *Turn) Add(newTurn Turn) bool {
	// if newTurn has more than one turn it is impossible to add
	if newTurn.prev != nil {
		return false
	}
	// if turn is empty
	if turn == nil {
		*turn = newTurn
		return true
	}

	cp := turn.Copy()

	*turn = newTurn

	turn.prev = cp

	return true
}

func (t Turn) Copy() *Turn {
	if t == (Turn{}) {
		return nil
	}
	return &t
}

func (t *Turn) String() string {
	var s string

	if t.prev != nil {
		s += t.prev.String() + " -> "
	}

	s += t.White.String() + "+" + t.Black.String()

	return s
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
