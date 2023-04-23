package main

import (
	"github.com/twanies/weavebox"
	"net/http"
)

func main() {
	app := weavebox.New()
	users := app.Box("/users")
	users.Get("/", handleGetUsers)
	users.Get("/:id", handleGetUserById)
}
func handleGetUsers(ctx *weavebox.Context) error {
	return ctx.Text(http.StatusOK, "hello from 2015")
}
func handleGetUserById(ctx *weavebox.Context) error {
	id := ctx.Param("id")
	return ctx.Text(http.StatusOK, id)
}
