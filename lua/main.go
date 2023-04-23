package main

import (
	"github.com/Shopify/go-lua"
	"log"
)

var script = `
bestBid = bestBid()
bestAsk = bestAsk()
spread = math.abs(bestBid-bestAsk);
print(spread);
`

func main() {
	l := lua.NewState()
	lua.OpenLibraries(l)
	registerBestBid(l)
	registerBestAsk(l)
	if err := lua.DoString(l, script); err != nil {
		log.Fatal(err)
	}
}
func registerBestBid(l *lua.State) {
	l.Register("bestBid", func(state *lua.State) int {
		l.PushInteger(10000)
		return 1
	})
}
func registerBestAsk(l *lua.State) {
	l.Register("bestAsk", func(state *lua.State) int {
		l.PushInteger(3333)
		return 1
	})
}
