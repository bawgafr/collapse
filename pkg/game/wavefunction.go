package game

import "github.com/hajimehoshi/ebiten/v2"

type WaveFunction struct {
	Id               int
	Name             string
	allowsNeighbours []int
	Img              *ebiten.Image
}

// func NewCollapseSet(game CollapseGame) []WaveFunction {

// 	mountain := WaveFunction{Name: "Mountain"}

// 	c := []WaveFunction{}

// 	return c
// }
