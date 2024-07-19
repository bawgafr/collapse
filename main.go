package main

import (
	"embed"
	"log"

	"github.com/bawgaft/collapse/pkg/game"
	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed images/*
var embededStatic embed.FS

func main() {
	game, err := game.NewGame(embededStatic)


	if err != nil {
		log.Fatal("error initialising game :", err)
	}
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowSize(640, 640)
	ebiten.SetWindowTitle("Collapse")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}

}
