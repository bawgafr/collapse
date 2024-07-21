package main

import (
	"embed"
	"log"

	"github.com/bawgaft/collapse/pkg/game"
	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed static/images/* static/rules/*
var embededStatic embed.FS

func main() {
	game := game.NewGame(embededStatic)

	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowSize(640, 640)
	ebiten.SetWindowTitle("Collapse")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}

}
