package controllers

import (
	"github.com/revel/revel"
)

// App -- main application controller
type App struct {
	Controller
}

// Index -- home
func (c App) Index() revel.Result {
	return c.Render()
}

// About -- about
func (c App) About() revel.Result {
	return c.Render()
}
