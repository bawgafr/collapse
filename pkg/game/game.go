package game

import (
	"embed"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type CollapseGame struct {
	fs       embed.FS
	Mountain *WaveFunction
	Plains   *WaveFunction
	Forest   *WaveFunction
	Swamp    *WaveFunction
	Beach    *WaveFunction
	Sea      *WaveFunction
}

const (
	cardSize     = 128
	resizeFactor = 0.25
)

func (g CollapseGame) Update() error {
	return nil
}

// draw a card to the grid
//
//	x : is the horizontal position in the grid that the image is to be displayed at
//	y : is the vertical position in the grid that the image is to be display at
//	w : is the card to be drawn at this position
//	screen : is the image that its being adde to
func drawCard(x, y int, w *WaveFunction, screen *ebiten.Image) {
	drawOptions := &ebiten.DrawImageOptions{}
	newSize := cardSize * resizeFactor
	xPos := float64(x) * newSize
	yPos := float64(y) * newSize

	drawOptions.GeoM.Scale(resizeFactor, resizeFactor)

	drawOptions.GeoM.Translate(float64(xPos), float64(yPos))
	screen.DrawImage(w.Img, drawOptions)

}

func (g CollapseGame) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, world")

	drawCard(0, 0, g.Mountain, screen)
	drawCard(1, 0, g.Plains, screen)
	drawCard(2, 0, g.Forest, screen)
	drawCard(0, 1, g.Swamp, screen)
	drawCard(1, 1, g.Beach, screen)
	drawCard(2, 1, g.Sea, screen)

}

func (g CollapseGame) Layout(outsideWidth, outsideHeight int) (screenWidth, ScreenHeight int) {
	return 640, 640
}

func NewGame(fs embed.FS) (CollapseGame, error) {
	game := CollapseGame{}
	game.fs = fs
	err := game.init()

	return game, err
}

func (g *CollapseGame) init() error {
	var err error

	// set up the cards for the collapse. Alowes Neighbours are the ids of cards that are
	// allowed to appear in the 4 cardinal positions around the card
	g.Mountain = &WaveFunction{Id: 1, allowsNeighbours: []int{1, 2}}
	g.Plains = &WaveFunction{Id: 2, allowsNeighbours: []int{2, 1, 3, 4, 5}}
	g.Forest = &WaveFunction{Id: 3, allowsNeighbours: []int{2, 4, 5}}
	g.Swamp = &WaveFunction{Id: 4, allowsNeighbours: []int{4, 2, 3}}
	g.Beach = &WaveFunction{Id: 5, allowsNeighbours: []int{5, 2, 3, 6}}
	g.Sea = &WaveFunction{Id: 6, allowsNeighbours: []int{6, 5}}

	// load all of the images
	if g.Mountain.Img, _, err = ebitenutil.NewImageFromFileSystem(g.fs, "images/mountain.png"); err != nil {
		return err
	}
	if g.Forest.Img, _, err = ebitenutil.NewImageFromFileSystem(g.fs, "images/forrest.png"); err != nil {
		return err
	}
	if g.Plains.Img, _, err = ebitenutil.NewImageFromFileSystem(g.fs, "images/plains.png"); err != nil {
		return err
	}
	if g.Swamp.Img, _, err = ebitenutil.NewImageFromFileSystem(g.fs, "images/swamp.png"); err != nil {
		return err
	}
	if g.Beach.Img, _, err = ebitenutil.NewImageFromFileSystem(g.fs, "images/beach.png"); err != nil {
		return err
	}
	if g.Sea.Img, _, err = ebitenutil.NewImageFromFileSystem(g.fs, "images/sea.png"); err != nil {
		return err
	}

	return nil
}
