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
	r.Error = err + "\n"
}

func (r *Render) EmptyErrorLine() {
	r.Error = "___________________\n"
}

func (r Render) Print(body string) {
	fmt.Print(body)
	fmt.Print(">\n")
	fmt.Print(r, "\n")
}

func (r Render) PrintBoard(board session.Board) {
	r.Print(r.StringBoard(board))
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
		r.Print(r.Menu)
	}
}

func getRandomSide() session.Side {
	rand.Seed(rand.Int63())
	if rand.Intn(2) == 0 {
		return session.WhiteSide
	}
	return session.BlackSide
}

func (r *Render) Run() {
	for {
		clear()

		r.PrintMenu()

		input := r.Scan()

		if input == "1" {
			r.Chess()
		}
		if input == "2" {
			r.FEN()
		}
		if input == "3" {
			r.PGN()
		}
		if input == "4" {
			r.TestGetPiece()
		}
		if input == "5" {
			r.TestPly()
		}
		if input == "q" {
			break
		}
	}
}

// Scan reads a line from stdin. NOTE: However, if err is not nil, it is not stop upper function
func (r *Render) Scan() string {
	fmt.Print("\033[11;2H")
	input := ""
	var err error

	// get input without fmt.Scanln()
	for {
		c := make([]byte, 1)
		_, err = os.Stdin.Read(c)
		if err != nil {
			return ""
		}
		if c[0] == '\n' {
			break
		}
		input += string(c)
	}

	if err != nil {
		r.ErrorLine(err.Error())
	}

	return input
}

// colour terminal, clear screen and move cursor to top left
func clear() {
	fmt.Print("\x1b[48;5;0m\x1b[38;5;231m\033[H\033[2J\033[1;1H")
}

// string representation of a board
func (r Render) StringBoard(b session.Board) string {
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
