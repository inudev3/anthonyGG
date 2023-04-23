package main

type Storage interface {
	Get(int) *User
}
