package app

import (
	log "github.com/sirupsen/logrus"
	"github.com/vehsamrak/gochess/internal/chess"
	"github.com/vehsamrak/gochess/internal/logger"
	"github.com/vehsamrak/gochess/internal/render/consoleRender"
)

type Game struct {
}

func (game Game) Start() {
	log.SetFormatter(&logger.TextFormatter{})
	render := consoleRender.ConsoleRender{}

	log.Info("Game started")

	x := 8
	y := 8
	board := chess.Board{}.Create(x, y)

	render.Render(*board)

	log.Info("Game finished")
}
