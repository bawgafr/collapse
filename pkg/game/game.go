package game

import (
	"embed"

	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	cardSize     = 128
	resizeFactor = 0.5
	gameSize     = 10
)

type CollapseGame struct {
	fs    embed.FS
	cards map[int]WaveFunction
	board [][]WaveFunction
}

// type Cards struct {
// 	Mountain *WaveFunction
// 	Plains   *WaveFunction
// 	Forest   *WaveFunction
// 	Swamp    *WaveFunction
// 	Beach    *WaveFunction
// 	Sea      *WaveFunction
// 	Blank    *WaveFunction
// 	mapped   map[int]WaveFunction
// }

// func (c Cards) enMappen() {

// }

// work out which cards are allowed in position x,y by getting the allowedNeighbours of cards in the cardinal positions. If edge is found, nothing is added
func (g CollapseGame) getEntropy(x, y int) []int {

	entropy := make(map[int]struct{})

	if x > 0 {
		for a := range g.board[x-1][y].allowsNeighbours {
			entropy[a] = struct{}{}
		}
	}

	if x < gameSize-1 {
		for a := range g.board[x+1][y].allowsNeighbours {
			entropy[a] = struct{}{}
		}
	}

	if y > 0 {
		for a := range g.board[x][y-1].allowsNeighbours {
			entropy[a] = struct{}{}
		}
	}

	if y < gameSize-1 {
		for a := range g.board[x][y+1].allowsNeighbours {
			entropy[a] = struct{}{}
		}

	}

	e := make([]int, len(entropy))
	i := 0
	for k := range entropy {
		e[i] = k
		i++
	}

	return e
}

// look at the board and find the places with the lowest entropy -- that is the fewest possible number of cards
func (g CollapseGame) unroll() []unrolledBoard {
	var unrolled []unrolledBoard

	smallestEntropy := 100000 // large number
	for i, _ := range g.board {
		for j, w := range g.board[i] {
			if w.Id == 0 {
				entropy := g.getEntropy(i, j)
				l := len(entropy)
				if l < smallestEntropy {
					unrolled = []unrolledBoard{}
					unrolled = append(unrolled, unrolledBoard{x: i, y: j, e: entropy})
				}
				if l == smallestEntropy {
					unrolled = append(unrolled, unrolledBoard{x: i, y: j, e: entropy})
				}
			}
		}
	}

	return unrolled
}

// draw a card to the grid
//
//	x : is the horizontal position in the grid that the image is to be displayed at
//	y : is the vertical position in the grid that the image is to be display at
//	w : is the card to be drawn at this position
//	screen : is the image that its being adde to
func drawCard(x, y int, w WaveFunction, screen *ebiten.Image) {
	if w.Img == nil {
		return
	}
	drawOptions := &ebiten.DrawImageOptions{}
	newSize := cardSize * resizeFactor
	xPos := float64(x) * newSize
	yPos := float64(y) * newSize

	drawOptions.GeoM.Scale(resizeFactor, resizeFactor)

	drawOptions.GeoM.Translate(float64(xPos), float64(yPos))
	screen.DrawImage(w.Img, drawOptions)

}

func NewGame(fs embed.FS) CollapseGame {
	game := CollapseGame{}
	game.fs = fs
	game.init()

	return game
}

func (g *CollapseGame) addCard(id int, name, filename string, allowedNeighbours []int) {
	wf := NewWaveFunction(id, name, filename, allowedNeighbours, g.fs)
	g.cards[id] = wf
}

func (g *CollapseGame) init() {

	a := make([][]WaveFunction, gameSize)
	for i := range a {
		a[i] = make([]WaveFunction, gameSize)
	}
	g.board = a
	g.cards = make(map[int]WaveFunction)
	// set up the cards for the collapse. Alowes Neighbours are the ids of cards that are
	// allowed to appear in the 4 cardinal positions around the card
	g.addCard(0, "blank", "", []int{1, 2, 3, 4, 5, 6})
	g.addCard(1, "Mountain", "images/mountain.png", []int{1, 2})
	g.addCard(2, "Plains", "images/forrest.png", []int{2, 1, 3, 4, 5})
	g.addCard(3, "Forest", "images/plains.png", []int{2, 4, 5})
	g.addCard(4, "Swamp", "images/swamp.png", []int{4, 2, 3})
	g.addCard(5, "Beach", "images/beach.png", []int{5, 2, 3, 6})
	g.addCard(6, "Sea", "images/sea.png", []int{6, 5})

	g.board[0][0] = g.cards[0]

	g.board[1][1] = g.cards[3]
	g.board[9][9] = g.cards[1]

	//	unrolled := g.unroll()

	return
}

func (g *CollapseGame) evolveBoard() {
	// unrolled should contain the cells with the smallest entropy and only those cells
	unrolled := g.unroll()

	if len(unrolled) == 0 {
		return
	}

	//get a random cell
	card := unrolled[rand.Intn(len(unrolled))]

	if len(card.e) == 0 {
		return
	}

	// pick one of the possible neighbours
	wf := card.e[rand.Intn(len(card.e))]

	g.board[card.x][card.y] = g.cards[wf]

}
