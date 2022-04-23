package main

import "fmt"

type character struct {
	dexterity int
	name      string
}

type arena struct {
	party    []*character
	ennemies []*character
}

type turnOrder struct {
	characters []*character
}

func main() {
	e1 := &character{name: "Ennemy 1", dexterity: 1}
	e2 := &character{name: "Ennemy 2", dexterity: 2}
	e3 := &character{name: "Ennemy 3", dexterity: 5}

	c1 := &character{name: "Character 1", dexterity: 2}
	c2 := &character{name: "Character 2", dexterity: 6}

	a := arena{party: []*character{c1, c2}, ennemies: []*character{e1, e2, e3}}
	fmt.Println("Arena", a)

}
