package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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
	ebitenutil.DebugPrint(screen, "Hello, world")

	// drawCard(0, 0, g.cards[1], screen)
	// drawCard(1, 0, g.cards[2], screen)
	// drawCard(2, 0, g.cards[3], screen)
	// drawCard(0, 1, g.cards[4], screen)
	// drawCard(1, 1, g.cards[5], screen)
	// drawCard(2, 1, g.cards[6], screen)

	for x := 0; x < gameSize; x++ {
		for y := 0; y < gameSize; y++ {
			drawCard(x, y, g.board[x][y], screen)
		}

	}

}

func (g CollapseGame) Layout(outsideWidth, outsideHeight int) (screenWidth, ScreenHeight int) {
	return 640, 640
}
