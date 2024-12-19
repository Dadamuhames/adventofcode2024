package main

type guard struct {
	posX int
	posY int
	dir  string
}

func (g *guard) up() (int, int) {
	return g.posX - 1, g.posY
}

func (g *guard) right() (int, int) {
	return g.posX, g.posY + 1
}

func (g *guard) down() (int, int) {
	return g.posX + 1, g.posY
}

func (g *guard) left() (int, int) {
	return g.posX, g.posY - 1
}

func (g guard) next() (int, int) {
	switch g.dir {
	case "^":
		return g.up()

	case ">":
		return g.right()

	case "v":
		return g.down()

	case "<":
		return g.left()
	}

	return 0, 0
}

func (g *guard) changeDir() {
	switch g.dir {
	case "^":
		g.dir = ">"

	case ">":
		g.dir = "v"

	case "v":
		g.dir = "<"

	case "<":
		g.dir = "^"
	}
}
