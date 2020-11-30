package consoleRender

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/vehsamrak/gochess/internal/chess"
)

type ConsoleRender struct {
}

func (render *ConsoleRender) Render(board chess.Board) {
	x, y := board.Size()

	for i := 1; i <= y; i++ {
		var rowString string
		for j := 1; j <= x; j++ {
			figure := board.Figure(j, i)
			if figure == nil {
				rowString += "[ ]"
			} else {
				rowString += fmt.Sprintf("[%s]", figure.Name())
			}
		}

		log.Infof(rowString)
	}

}
