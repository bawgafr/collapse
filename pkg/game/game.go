package game

import (
	"embed"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type CollapseGame struct {
	fs            embed.FS
	MountainImage *ebiten.Image
	PlainsImage   *ebiten.Image
	ForrestImage  *ebiten.Image
	SwampImage    *ebiten.Image
	BeachImage    *ebiten.Image
	SeaImage      *ebiten.Image
}

var cardSize = 128

func (g CollapseGame) Update() error {
	return nil
}

func drawCard(x, y int, img *ebiten.Image, screen *ebiten.Image) {
	drawOptions := &ebiten.DrawImageOptions{}
	drawOptions.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(img, drawOptions)

}

func (g CollapseGame) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, world")

	drawCard(0, 0, g.MountainImage, screen)
	drawCard(cardSize, 0, g.PlainsImage, screen)
	drawCard(cardSize*2, 0, g.ForrestImage, screen)
	drawCard(0, cardSize, g.SwampImage, screen)
	drawCard(cardSize, cardSize, g.BeachImage, screen)
	drawCard(cardSize*2, cardSize, g.SeaImage, screen)

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

	if g.MountainImage, _, err = ebitenutil.NewImageFromFileSystem(g.fs, "images/mountain.png"); err != nil {
		return err
	}

	if g.ForrestImage, _, err = ebitenutil.NewImageFromFileSystem(g.fs, "images/forrest.png"); err != nil {
		return err
	}

	if g.PlainsImage, _, err = ebitenutil.NewImageFromFileSystem(g.fs, "images/plains.png"); err != nil {
		return err
	}

	if g.SwampImage, _, err = ebitenutil.NewImageFromFileSystem(g.fs, "images/swamp.png"); err != nil {
		return err
	}

	if g.BeachImage, _, err = ebitenutil.NewImageFromFileSystem(g.fs, "images/beach.png"); err != nil {
		return err
	}

	if g.SeaImage, _, err = ebitenutil.NewImageFromFileSystem(g.fs, "images/sea.png"); err != nil {
		return err
	}

	return nil
}
