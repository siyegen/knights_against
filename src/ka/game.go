package main

type Game struct {
	level *Level
	char  *Character
}

type Level struct {
	xSize int
	ySize int
}

type Pos [2]int

type Character struct {
	pos Pos
}

func CreateGame(xSize, ySize, xStart, yStart int) *Game {
	var start Pos
	start[0] = xStart
	start[1] = yStart
	return &Game{
		level: &Level{xSize, ySize},
		char:  &Character{start},
	}
}
