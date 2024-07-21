package game

import (
	"embed"
	"log"
	"slices"

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

func intersection(s1, s2 []int) []int {
	if len(s1) > len(s2) {
		s1, s2 = s2, s1
	}

	intersection := []int{}

	for _, v := range s1 {
		if slices.Contains(s2, v) {
			intersection = append(intersection, v)
		}
	}
	return intersection
}

func intersectionM(s1, s2 map[int]struct{}) map[int]struct{} {
	sIntersecting := make(map[int]struct{})

	if len(s1) > len(s2) {
		s1, s2 = s2, s1
	}

	for k, _ := range s1 {
		_, intersecting := s2[k]
		if intersecting {
			sIntersecting[k] = struct{}{}
		}
	}

	return sIntersecting
}

// work out which cards are allowed in position x,y by getting the allowedNeighbours of cards
// in the cardinal positions and finding an intersection of the available neighbours.
// If edge is found, intersection is not made
func (g CollapseGame) getEntropy(x, y int) []int {
	// needs to be the intersection of the available....

	entropy := g.cards[0].allowsNeighbours

	if x > 0 {
		cell := g.board[x-1][y]
		entropy = intersection(entropy, cell.allowsNeighbours)
	}

	if x < gameSize-1 {

		cell := g.board[x+1][y]
		entropy = intersection(entropy, cell.allowsNeighbours)
	}

	if y > 0 {
		cell := g.board[x][y-1]
		entropy = intersection(entropy, cell.allowsNeighbours)

	}

	if y < gameSize-1 {
		cell := g.board[x][y+1]
		entropy = intersection(entropy, cell.allowsNeighbours)

	}

	return entropy
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
					smallestEntropy = l
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

	g.cards = make(map[int]WaveFunction)
	// set up the cards for the collapse. Alowes Neighbours are the ids of cards that are
	// allowed to appear in the 4 cardinal positions around the card
	g.addCard(0, "blank", "", []int{1, 2, 3, 4, 5, 6})
	g.addCard(1, "Mountain", "static/images/mountain.png", []int{1, 2})
	g.addCard(2, "Plains", "static/images/plains.png", []int{2, 1, 3, 4, 5})
	g.addCard(3, "Forest", "static/images/forrest.png", []int{2, 4, 5})
	g.addCard(4, "Swamp", "static/images/swamp.png", []int{4, 2, 3})
	g.addCard(5, "Beach", "static/images/beach.png", []int{5, 2, 3, 6})
	g.addCard(6, "Sea", "static/images/sea.png", []int{6, 5})

	a := make([][]WaveFunction, gameSize, gameSize)
	for i := range a {
		for j := 0; j < gameSize; j++ {
			a[i] = append(a[i], g.cards[0])
		}
	}
	g.board = a

	// seed a card in there already
	g.board[1][0] = g.cards[3]
	return
}

func (g *CollapseGame) evolveBoard() bool {
	// unrolled should contain the cells with the smallest entropy and only those cells
	unrolled := g.unroll()
	if len(unrolled) == 0 {
		log.Println("unrolled == 0")
		return false
	}

	//get a random cell
	c := rand.Intn(len(unrolled))
	card := unrolled[c]

	if len(card.e) == 0 {
		log.Println("entropy == 0")
		return false
	}

	// pick one of the possible neighbours
	wf := card.e[rand.Intn(len(card.e))]

	// set the randomly picked card with the randomly picked possible neighbour
	g.board[card.x][card.y] = g.cards[wf]

	return true
}
