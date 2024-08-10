package game

import (
	"bufio"
	"encoding/json"
	"io/fs"
	"log"
	"slices"

	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

// const (
// 	cardSize     = 128
// 	resizeFactor = 0.5
// 	gameSize     = 10
// )

type CollapseGame struct {
	fs           fs.FS
	cardSize     int
	resizeFactor float64
	boardSize    int
	cards        map[int]WaveFunction
	board        [][]WaveFunction
	Randomiser   int
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

// work out which cards are allowed in position x,y by getting the allowedNeighbours of cards
// in the cardinal positions and finding an intersection of the available neighbours.
// If edge is found, intersection is not made
func (g CollapseGame) getEntropy(x, y int) []int {
	// needs to be the intersection of the available....

	entropy := g.cards[0].allowedNeighbours

	if x > 0 {
		cell := g.board[x-1][y]
		entropy = intersection(entropy, cell.allowedNeighbours)
	}

	if x < g.boardSize-1 {

		cell := g.board[x+1][y]
		entropy = intersection(entropy, cell.allowedNeighbours)
	}

	if y > 0 {
		cell := g.board[x][y-1]
		entropy = intersection(entropy, cell.allowedNeighbours)

	}

	if y < g.boardSize-1 {
		cell := g.board[x][y+1]
		entropy = intersection(entropy, cell.allowedNeighbours)

	}

	return entropy
}

// look at the board and find the places with the lowest entropy -- that is the fewest possible number of cards
func (g CollapseGame) unroll() []unrolledBoard {
	var unrolled []unrolledBoard

	smallestEntropy := 100000 // large number

	for i := range g.board {
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
func drawCard(x, y, cardSize int, resizeFactor float64, w WaveFunction, screen *ebiten.Image) {
	if w.Img == nil {
		return
	}
	drawOptions := &ebiten.DrawImageOptions{}
	newSize := float64(cardSize) * resizeFactor
	xPos := float64(x) * newSize
	yPos := float64(y) * newSize

	drawOptions.GeoM.Scale(resizeFactor, resizeFactor)

	drawOptions.GeoM.Translate(float64(xPos), float64(yPos))
	screen.DrawImage(w.Img, drawOptions)

}

func NewGame(fs fs.FS) CollapseGame {
	game := CollapseGame{}
	game.fs = fs

	// load the json from static/rules/rules.json
	rules, err := readRulesFromDisk("static/rules/rules.json", fs)
	if err != nil {
		log.Fatal("Unable to load Game Rules", err)
	}

	log.Println("rules", rules)

	game.cardSize = rules.ImageSize
	game.resizeFactor = rules.ResizeFactor
	game.boardSize = rules.BoardSize
	game.Randomiser = rules.Randomiser

	game.cards = make(map[int]WaveFunction)

	// create a blank card using all of the ids for allowed neighbours
	blank := cardRules{Id: 0, Name: "blank", Filename: ""}
	bAN := make([]int, len(rules.Cards))
	for _, c := range rules.Cards {
		bAN = append(bAN, c.Id)
	}
	blank.AllowedNeighbours = bAN
	rules.Cards = append(rules.Cards, blank)
	game.addCards(rules.Cards)

	// initialise the board with the blank card
	a := make([][]WaveFunction, game.boardSize)
	for i := range a {
		for j := 0; j < game.boardSize; j++ {
			a[i] = append(a[i], game.cards[0])
		}
	}
	game.board = a

	// seed a card in the board
	for _, seed := range rules.Seeds {
		game.board[seed.X][seed.Y] = game.cards[seed.Card]
	}

	return game
}

func (g *CollapseGame) addCards(cardRules []cardRules) {
	for _, c := range cardRules {
		if c.Filename != "" {
			c.Filename = "static/" + c.Filename
		}
		wf := NewWaveFunction(c, g.fs)
		g.cards[c.Id] = wf
	}
}

func (g *CollapseGame) evolveBoard() bool {
	// unrolled should contain the cells with the smallest entropy and only those cells
	// if the smallest amount of entropy is 0 we're either finished or backed into a corner.
	// assume we've finished for now.
	unrolled := g.unroll()
	if len(unrolled) == 0 {
		log.Println("unrolled == 0")
		return false
	}

	// select a location on the board to put the new card
	// it will be a random selection from the unrolled set
	// which contains all of the cells with the smallest entropy
	c := rand.Intn(len(unrolled))
	card := unrolled[c]

	// if the entropy of the card is 0 then we have a similar case to earlier
	// where it is either the end, or we've reached a point where a card cannot be placed
	if len(card.e) == 0 {
		log.Println("entropy == 0")
		return false
	}

	// using the list of available neighbours that the unrolledBoard object contains
	// select one of the card types at random to place in the pre-selected location

	// wf := card.e[rand.Intn(len(card.e))]
	newCardsIndex := g.randomCard(card)

	// set the randomly picked card with the randomly picked possible neighbour
	g.board[card.x][card.y] = g.cards[newCardsIndex]

	return true
}

func readRulesFromDisk(filename string, fs fs.FS) (GameRules, error) {
	f, err := fs.Open(filename)
	if err != nil {
		return GameRules{}, err
	}
	defer f.Close()

	var jsonString string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		jsonString += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return GameRules{}, err
	}

	rules := GameRules{}
	err = json.Unmarshal([]byte(jsonString), &rules)
	if err != nil {
		return GameRules{}, err
	}

	return rules, nil

}
