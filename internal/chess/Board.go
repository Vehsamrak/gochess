package chess

type Board struct {
	x     int
	y     int
	cells map[int]map[int]*Cell
}

func (Board) Create(x int, y int) *Board {
	cells := make(map[int]map[int]*Cell, y)
	for i := 1; i <= y; i++ {
		cells[i] = make(map[int]*Cell, x)
		for j := 1; j <= y; j++ {
			cells[i][j] = &Cell{}
		}
	}

	board := &Board{x: x, y: y, cells: cells}
	board.addStandardFigures()

	return board
}

func (board *Board) AddFigure(figure *Figure, x int, y int) {
	board.cells[x][y].Figure = figure
}

func (board *Board) Size() (x int, y int) {
	return board.x, board.y
}

// Cell retrieves Figure in cell, or nil if it is empty
func (board *Board) Figure(x int, y int) *Figure {
	cell := board.cells[x][y]

	if cell == nil {
		return nil
	}

	return cell.Figure
}

func (board *Board) addStandardFigures() {
	board.AddFigure(Figure{}.Create(Rook, White), 1, 8)
	board.AddFigure(Figure{}.Create(Knight, White), 2, 8)
	board.AddFigure(Figure{}.Create(Bishop, White), 3, 8)
	board.AddFigure(Figure{}.Create(Queen, White), 4, 8)
	board.AddFigure(Figure{}.Create(King, White), 5, 8)
	board.AddFigure(Figure{}.Create(Bishop, White), 6, 8)
	board.AddFigure(Figure{}.Create(Knight, White), 7, 8)
	board.AddFigure(Figure{}.Create(Rook, White), 8, 8)
	board.AddFigure(Figure{}.Create(Pawn, White), 1, 7)
	board.AddFigure(Figure{}.Create(Pawn, White), 2, 7)
	board.AddFigure(Figure{}.Create(Pawn, White), 3, 7)
	board.AddFigure(Figure{}.Create(Pawn, White), 4, 7)
	board.AddFigure(Figure{}.Create(Pawn, White), 5, 7)
	board.AddFigure(Figure{}.Create(Pawn, White), 6, 7)
	board.AddFigure(Figure{}.Create(Pawn, White), 7, 7)
	board.AddFigure(Figure{}.Create(Pawn, White), 8, 7)

	board.AddFigure(Figure{}.Create(Rook, Black), 1, 1)
	board.AddFigure(Figure{}.Create(Knight, Black), 2, 1)
	board.AddFigure(Figure{}.Create(Bishop, Black), 3, 1)
	board.AddFigure(Figure{}.Create(Queen, Black), 4, 1)
	board.AddFigure(Figure{}.Create(King, Black), 5, 1)
	board.AddFigure(Figure{}.Create(Bishop, Black), 6, 1)
	board.AddFigure(Figure{}.Create(Knight, Black), 7, 1)
	board.AddFigure(Figure{}.Create(Rook, Black), 8, 1)
	board.AddFigure(Figure{}.Create(Pawn, Black), 1, 2)
	board.AddFigure(Figure{}.Create(Pawn, Black), 2, 2)
	board.AddFigure(Figure{}.Create(Pawn, Black), 3, 2)
	board.AddFigure(Figure{}.Create(Pawn, Black), 4, 2)
	board.AddFigure(Figure{}.Create(Pawn, Black), 5, 2)
	board.AddFigure(Figure{}.Create(Pawn, Black), 6, 2)
	board.AddFigure(Figure{}.Create(Pawn, Black), 7, 2)
	board.AddFigure(Figure{}.Create(Pawn, Black), 8, 2)
}
