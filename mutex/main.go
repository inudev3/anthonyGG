package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type Player struct {
	mu     sync.RWMutex
	health int32
}

func NewPlayer() *Player {
	return &Player{health: 100}
}
func startUILoop(p *Player) {
	ticker := time.NewTicker(time.Second)
	for {

		fmt.Printf("player health: %d\n", p.getHealth())
		<-ticker.C
	}
}
func (p *Player) takeDamage(val int) {
	health := p.getHealth()
	atomic.StoreInt32(&p.health, int32(health-val))
}
func (p *Player) getHealth() int {
	return int(atomic.LoadInt32(&p.health))
}
func startGameLoop(p *Player) {
	ticker := time.NewTicker(time.Millisecond * 300)
	for {
		p.takeDamage(rand.Intn(40))
		if p.getHealth() <= 0 {
			fmt.Println("GAME OVER")
			break
		}
		<-ticker.C
	}
}
func main() {

}
