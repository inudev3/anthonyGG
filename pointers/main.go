package main

import "fmt"

type Player struct {
	health int
}

func takeDamage(player *Player) {
	fmt.Println("damaged!")
	dmg := 10
	player.health -= dmg
}

func main() {
	player := Player{health: 100}
	fmt.Printf("before explosion: %d", player.health)
	takeDamage(&player)
	fmt.Printf("after explosion: %d", player.health)
}
