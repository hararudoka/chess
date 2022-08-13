package main

import "github.com/hararudoka/chess/internal/render"

func main() {
	r, err := render.New()
	if err != nil {
		panic(err)
	}
	r.Run()
}
