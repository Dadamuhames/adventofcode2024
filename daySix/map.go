package main

type taskMap struct {
	g      guard
	data   [][]int
	exited bool
}

func (m taskMap) isObstacle(x int, y int) bool {
	return m.data[x][y] == 1
}

func (m taskMap) isExit(x int, y int) bool {
	dataLen := len(m.data)
	rowLen := len(m.data[dataLen-1])

	return x == dataLen || x == -1 || y == rowLen || y == -1
}

func (m *taskMap) move() {
	nextX, nextY := m.g.next()

	if m.isExit(nextX, nextY) {
		m.exited = true
	} else if m.isObstacle(nextX, nextY) {
		m.g.changeDir()
	} else {
		m.g.posX = nextX
		m.g.posY = nextY
	}
}
