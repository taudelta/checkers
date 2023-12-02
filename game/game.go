package game

import (
	"log"

	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/pixelgl"
)

type Game struct {
}

func (g Game) Run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Checkers",
		Bounds: pixel.R(0, 0, 1024, 768),
	}
	log.Println("run game")
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	for !win.Closed() {
		win.Update()
	}
	log.Println("stop game")
}

var game Game

func Start() {
	pixelgl.Run(game.Run)
}
