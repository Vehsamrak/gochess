package render

import "github.com/vehsamrak/gochess/internal/chess"

type Renderer interface {
	Render(board chess.Board)
}
