package game

import (
	"fmt"
	"io/fs"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	naive  = 1
	chance = 2
)

type WaveFunction struct {
	Id                int
	Name              string
	allowedNeighbours []int
	Img               *ebiten.Image
	chance            float64
}

type GameRules struct {
	ImageSize    int
	BoardSize    int
	ResizeFactor float64
	Cards        []cardRules
	Seeds        []seedRules
	RandSeed     int64
	Randomiser   int
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
	Chance            float64
}

func NewWaveFunction(cardRule cardRules, fs fs.FS) WaveFunction {
	wf := WaveFunction{Id: cardRule.Id, Name: cardRule.Name, allowedNeighbours: cardRule.AllowedNeighbours, chance: cardRule.Chance}
	var err error
	if cardRule.Filename != "" {
		if wf.Img, _, err = ebitenutil.NewImageFromFileSystem(fs, cardRule.Filename); err != nil {
			log.Fatal("unable to load card image: ", err)
		}
	}
	return wf
}

type unrolledBoard struct {
	x, y int
	e    []int
}

// return an integer representing the id of the card that has been selected
func (g CollapseGame) randomCard(u unrolledBoard) int {

	switch g.Randomiser {
	// naive random assuming all cards equally likely
	case naive:
		return u.e[rand.Intn(len(u.e))]
	case chance:
		// do something clever -- remember cards are indexed from 1 not 0!
		// if u.e = []

		tot := 0.0
		boundary := []float64{}

		for _, index := range u.e {
			tot += g.cards[index].chance
			boundary = append(boundary, tot)
		}

		r := rand.Float64() * tot // 0->1

		for i, b := range boundary {
			if r <= b {
				fmt.Printf("returning %d for random (r=%f) %v %v\n", i, r, boundary, u.e)
				return i + 1
			}
		}
		fmt.Println("ooop")
	}
	fmt.Println("returning 1 for random")
	return 1
}
