package main

import "fmt"

type Player struct {
	*Position
}
type Position struct {
	x, y float64
}

func (p *Position) Move(x, y float64) {
	p.x += x
	p.y += y
}
func (p *Position) Teleport(x, y float64) {
	p.x = x
	p.y = y
}
func NewPlayer() *Player {
	return &Player{Position: &Position{}}
}

type Enemy struct {
	*SpecialPosition
}

func NewEnemy() *Enemy {
	return &Enemy{SpecialPosition: &SpecialPosition{}}
}
func main() {
	player := NewPlayer()
	//raidBoss := NewEnemy()

	fmt.Printf("%#+v\n", player.Position)
	player.Move(1, 3)
	fmt.Printf("%#+v\n", player.Position)
	player.Teleport(3, 5)
	fmt.Printf("%#+v\n", player.Position)
}

type SpecialPosition struct {
	Position
}

func (sp *SpecialPosition) MoveSpecial(x, y float64) {
	sp.x += x * x
	sp.y += y * y
}
