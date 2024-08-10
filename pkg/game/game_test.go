package game

import (
	"io/fs"
	"testing"
	"testing/fstest"
)

func getFs() fs.FS {
	fs := fstest.MapFS{
		"static/rules/rules.json": {
			Data: []byte(`{
	"imageSize":128,
	"boardSize":10,
	"resizeFactor": 0.5,
	"cards": [
		{"Id":1, "Name":"Mountain", "filename":"", "allowedNeighbours":[1, 2]},
		{"Id":2, "Name":"Plains",   "filename":"",   "allowedNeighbours":[2, 1, 3, 4, 5]},
		{"Id":3, "Name":"Forest",   "filename":"",  "allowedNeighbours":[2, 4, 5]},
		{"Id":4, "Name":"Swamp",    "filename":"",    "allowedNeighbours":[4, 2, 3]},
		{"Id":5, "Name":"Beach",    "filename":"",    "allowedNeighbours":[5, 2, 3, 6]},
		{"Id":6, "Name":"Sea",      "filename":"",      "allowedNeighbours":[6, 5]}
	],
	"seeds" : [
		{"x":3, "y":3, "card":3}
	]
}`)},
		"hello-world2.md": {Data: []byte("hola")},
	}

	return fs
}

func TestInitialise(t *testing.T) {
	t.Run("initialise a game", func(t *testing.T) {
		fs := getFs()
		game := NewGame(fs)

		if game.board[3][3].Id != 3 {
			t.Error("the 1,0 piece should be 3")
		}

		an := game.board[3][3].allowedNeighbours
		if len(an) != 3 {
			t.Error("Wrong allowed neighbours:", an)
		}

		if an[0] != 2 {
			t.Error("wrong first neighbour: ", an[0])
		}

	})

}

func TestSpeed(t *testing.T) {
	fs := getFs()
	g := NewGame(fs)

	for {
		if !g.evolveBoard() {
			break
		}

	}

}
