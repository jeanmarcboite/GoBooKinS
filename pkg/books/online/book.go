package online

import (
	"fmt"

	"github.com/jeanmarcboite/bookins/pkg/books/online/goodreads"
	"github.com/jeanmarcboite/bookins/pkg/books/online/openlibrary"
)

// Book -- book metadata
type Book struct {
	Title       string
	SubTitle    string
	Authors     []string
	Identifiers Identifiers
	URL         map[string]string
	Cover       map[string]string
	Subjects    []interface{}

	Inforigin   []interface{}
	Description string
	Metadata    Metadata
}

// Metadata -- metadata coming from www
type Metadata struct {
	// EPUB            epub.Metadata
	Goodreads   goodreads.Book
	Openlibrary openlibrary.Book
}

// Identifiers -- book identifiers
type Identifiers struct {
	ISBN10       string
	ISBN13       string
	LCCN         string
	Openlibrary  string
	Goodreads    string
	Librarything string
}

// New -- create a Book
func New(g goodreads.Book, o openlibrary.Book) Book {
	b := Book{
		Title: g.Title,
		Metadata: Metadata{
			Goodreads:   g,
			Openlibrary: o,
		},
	}

	b.Cover = make(map[string]string)
	if len(o.Data.Details.Covers[0]) > 0 {
		b.Cover["openlibrary"] = o.Data.Details.Covers[0]
	}
	if g.ImageURL != "" {
		b.Cover["goodreads"] = g.ImageURL
	}

	b.Authors = o.AuthorsName()
	b.URL = make(map[string]string)
	b.URL["preview"] = o.Data.PreviewURL
	b.URL["openlibrary"] = o.Data.InfoURL

	b.Identifiers.ISBN10 = g.ISBN
	b.Identifiers.ISBN13 = g.ISBN13
	b.Identifiers.Goodreads = g.ID
	b.Identifiers.Openlibrary = o.Data.Details.Key

	b.Description = fmt.Sprintf("goodreads: %v", g.Description)

	return b
}

// LookUpISBN -- lookup a work on goodreads and openlibrary, with isbn
func LookUpISBN(isbn string) (Book, error) {
	g, err := goodreads.LookUpISBN(isbn)

	if err != nil {
		return Book{}, err
	}
	o, err := openlibrary.LookUpISBN(isbn)

	return New(g.Books[0], o), nil
}

/*
// SearchTitle -- search for a work with a title
func SearchTitle(title string) (Book, error) {
	return Book{
		Metadata: {
			Goodreads:   goodreads.Search(title),
			Openlibrary: openlibrary.Search(title, ""),
		},
	}, nil
}

// SearchTitleAuthor -- search for a work with a title and an author
func SearchTitleAuthor(title string, author string) (Book, error) {
	return Book{
		Metadata: {
			Goodreads:   goodreads.Search(title),
			Openlibrary: openlibrary.Search(title, author),
		},
	}, nil
}
*/
