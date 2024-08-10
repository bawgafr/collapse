package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

var building = true

// enbitten function
func (g CollapseGame) Update() error {
	if building {
		building = g.evolveBoard()
	}

	return nil
}

// embitten function
func (g CollapseGame) Draw(screen *ebiten.Image) {

	for x := 0; x < g.boardSize; x++ {
		for y := 0; y < g.boardSize; y++ {
			drawCard(x, y, g.cardSize, g.resizeFactor, g.board[x][y], screen)
		}

	}

}

func (g CollapseGame) Layout(outsideWidth, outsideHeight int) (screenWidth, ScreenHeight int) {
	return 640, 640
}
