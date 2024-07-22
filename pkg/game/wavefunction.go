package game

import (
	"io/fs"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type WaveFunction struct {
	Id                int
	Name              string
	allowedNeighbours []int
	Img               *ebiten.Image
}

type GameRules struct {
	ImageSize    int
	BoardSize    int
	ResizeFactor float64
	Cards        []cardRules
	Seeds        []seedRules
}

type seedRules struct {
	X    int
	Y    int
	Card int
}

type cardRules struct {
	Id                int
	Name              string
	Filename          string
	AllowedNeighbours []int
}

func NewWaveFunction(id int, name, filename string, allowsNeighbours []int, fs fs.FS) WaveFunction {
	wf := WaveFunction{Id: id, Name: name, allowedNeighbours: allowsNeighbours}
	var err error
	if filename != "" {
		if wf.Img, _, err = ebitenutil.NewImageFromFileSystem(fs, filename); err != nil {
			log.Fatal("unable to load card image: ", err)
		}
	}
	return wf
}

type unrolledBoard struct {
	x, y int
	e    []int
}
