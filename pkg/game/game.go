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

var tiles []ebiten.Image

var cardSize = 128

func (g CollapseGame) Update() error {
	return nil
}

func (g CollapseGame) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, world")

	drawOptions := &ebiten.DrawImageOptions{}
	drawOptions.GeoM.Translate(0, 0)
	screen.DrawImage(g.MountainImage, drawOptions)

	drawOptions.GeoM.Translate(128, 0)
	screen.DrawImage(g.PlainsImage, drawOptions)

	drawOptions.GeoM.Translate(256, 0)
	screen.DrawImage(g.SwampImage, drawOptions)

	drawOptions.GeoM.Translate(384, 0)
	screen.DrawImage(g.ForrestImage, drawOptions)

	drawOptions.GeoM.Translate(0, 128)
	screen.DrawImage(g.BeachImage, drawOptions)

	drawOptions.GeoM.Translate(128, 128)
	screen.DrawImage(g.SeaImage, drawOptions)

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

	image, _, err := ebitenutil.NewImageFromFileSystem(g.fs, "images/mountain.png")
	if err != nil {
		return err
	}
	g.MountainImage = image

	image, _, err = ebitenutil.NewImageFromFileSystem(g.fs, "images/forrest.png")
	if err != nil {
		return err
	}
	g.ForrestImage = image

	image, _, err = ebitenutil.NewImageFromFileSystem(g.fs, "images/plains.png")
	if err != nil {
		return err
	}
	g.PlainsImage = image

	image, _, err = ebitenutil.NewImageFromFileSystem(g.fs, "images/swamp.png")
	if err != nil {
		return err
	}
	g.SwampImage = image

	image, _, err = ebitenutil.NewImageFromFileSystem(g.fs, "images/beach.png")
	if err != nil {
		return err
	}
	g.BeachImage = image

	image, _, err = ebitenutil.NewImageFromFileSystem(g.fs, "images/sea.png")
	if err != nil {
		return err
	}
	g.SeaImage = image

	tiles = []ebiten.Image{*g.BeachImage,
		*g.ForrestImage,
		*g.MountainImage,
		*g.PlainsImage,
		*g.SeaImage,
		*g.SwampImage,
	}

	return nil
}
