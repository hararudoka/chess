package render

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/hararudoka/chess/internal/session"
)

type Render struct {
	List  []string
	Error string

	State string

	Menu string
}

func (r Render) String() string {
	s := r.Error
	for _, e := range r.List {
		s += fmt.Sprint(e) + "\n"
	}
	return s
}

func (r *Render) Add(e string) {
	r.List = append(r.List, e)
}

func (r *Render) ErrorLine(err string) {
	r.Error = err
}

func (r *Render) EmptyErrorLine() {
	r.Error = "___________________"
}

func (r Render) Print(body string) {
	fmt.Print(body)
	fmt.Print(">\n")
	fmt.Print(r)
}

func New() (*Render, error) {
	main, err := os.ReadFile("menu.txt")
	if err != nil {
		return nil, err
	}

	menu := &Render{
		State: "menu",
		Menu:  string(main),
	}

	return menu, err
}

// PrintMenu prints 10 lines of a menu depending on the state of the game
func (r Render) PrintMenu() {
	if r.State == "menu" {
		fmt.Print(r.Menu)
	}
}

func getRandomSide() session.Side {
	rand.Seed(rand.Int63())
	if rand.Intn(2) == 0 {
		return session.White
	}
	return session.Black
}

func (r *Render) Run() {
	for {
		clear()
		r.PrintMenu()
		fmt.Print(">\n\033[11;2H\n")
		fmt.Print(r)

		r.EmptyErrorLine()

		input, err := r.Scan()
		if err != nil {
			r.ErrorLine(err.Error())
			continue
		}

		if input == "1" {
			r.Chess()
		}
		if input == "2" {
			r.FEN()
		}
		if input == "3" {
			r.PGN()
		}
		if input == "q" {
			break
		}
	}
}

func (r *Render) FEN() {
	currentBoard := session.Board{}
	for {
		clear()

		r.Print(r.RenderBoard(currentBoard))

		fen, err := r.Scan()
		if err != nil {
			r.ErrorLine(err.Error())
			continue
		}

		round, err := session.FENToRound(fen)
		if err != nil {
			r.ErrorLine(err.Error())
			continue
		}

		currentBoard = round.Board
	}
}

func (r *Render) PGN() {
	pgn := "*"
	for {
		clear()
		s, err := session.New(pgn, "chess")
		if err != nil {
			r.ErrorLine(err.Error())
			break
		}

		board := r.RenderBoard(s.Round.Board)
		r.Print(board)

		pgnFile, err := r.Scan()
		if err != nil {
			r.ErrorLine(err.Error())
			continue
		}

		pgnBytes, err := os.ReadFile(pgnFile)
		if err != nil {
			r.ErrorLine(err.Error())
			break
		}

		pgn = string(pgnBytes)
	}
}

func (r Render) Scan() (string, error) {
	fmt.Print("\033[11;2H")
	input := ""

	// get input without fmt.Scanln()
	for {
		c := make([]byte, 1)
		_, err := os.Stdin.Read(c)
		if err != nil {
			return "", err
		}
		if c[0] == '\n' {
			break
		}
		input += string(c)
	}

	return input, nil
}

// colour terminal, clear screen and move cursor to top left
func clear() {
	fmt.Print("\x1b[48;5;0m\x1b[38;5;231m\033[H\033[2J\033[1;1H")
}

// basic render of the board
func (r Render) RenderBoard(b session.Board) string {
	designs := map[string]string{
		"P": "♟",
		"R": "♜",
		"Q": "♛",
		"K": "♚",
		"B": "♝",
		"N": "♞",

		"p": "♙",
		"r": "♖",
		"q": "♕",
		"k": "♔",
		"b": "♗",
		"n": "♘",
	}
	s := "  A B C D E F G H\n"
	for i, row := range b {
		s += fmt.Sprint(8 - i)
		for _, piece := range row {
			if (piece == session.Piece{}) {
				s += "| "
			} else {
				s += "|" + designs[piece.Kind]
			}
		}
		s += "|" + fmt.Sprint(8-i) + "\n"
	}
	s += "  A B C D E F G H\n"
	return s
}

func (r *Render) Chess() {
	s, err := session.New("*", "chess")
	if err != nil {
		r.ErrorLine(err.Error())
		return
	}
	s.Round.SideOfPlayer = session.White
	for {
		clear()

		board := r.RenderBoard(s.Round.Board)
		r.Print(board)

		input, err := r.Scan()
		if err != nil {
			r.ErrorLine(err.Error())
			continue
		}

		s.Round.RawPly(input)
	}
}