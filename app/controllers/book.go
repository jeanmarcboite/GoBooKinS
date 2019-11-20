package controllers

import (
	"github.com/jeanmarcboite/bookins/pkg/books/online"
	"github.com/revel/revel"
)

// Book -- controller for Book display
type Book struct {
	Controller
}

// Details -- display book details
func (c Book) Details(id string) revel.Result {
	book, err := online.LookUpISBN(id)

	if err != nil {
		return c.Render(err)
	}

	title := book.Title
	data := c.SprintHTML(book)

	return c.Render(title, data)
}
