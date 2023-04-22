package main

import (
	"fmt"
	"net"
	"sync"
)

type Unit float32

const (
	Version Unit = 0.1 * iota
	Scalar
)

func Foo() int {
	var (
		x   = 100
		y   = 2
		foo = "foo"
	)
	fmt.Println(foo)
	return x + y
}

type Server struct {
	listenAddr string
	isRunning  bool
	peers      map[string]net.Conn
	peerLock   sync.Mutex
}
type Getter interface {
	Get()
}
type Putter interface {
	Put()
}
type Patcher interface {
	Patch()
}
type Deleter interface {
	Deleter()
}
type Storer interface {
	Getter
	Putter
	Deleter
	Patcher
}
type Vector struct {
	x, y int
}

func MustParseIntFromString(s string) (int, error) {
	//logic
	panic("oops")
	return 0, nil
}
