package openlibrary

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jeanmarcboite/bookins/pkg/books/online/net"
)

// LookUpISBN -- lookup a work on goodreads, with isbn
func LookUpISBN(isbn string) (Book, error) {
	url := fmt.Sprintf(net.Koanf.String("openlibrary.url.isbn"), isbn)
	olresponse, err := net.HTTPGet(url)
	response := strings.Replace(string(olresponse), "ISBN:9780061001116", "data", 1)

	if err != nil {
		return Book{}, err
	}

	var book Book
	json.Unmarshal([]byte(response), &book)

	return book, nil
}

// SearchTitle -- search for a work with a title
func SearchTitle(title string, author string) (Response, error) {
	w := title
	if idx := strings.IndexAny(w, "(-"); idx >= 0 {
		w = w[:idx]
	}
	plusWords := strings.Join(strings.Fields(w), "+")

	var url string
	if len(author) <= 0 {
		url = fmt.Sprintf(net.Koanf.String("openlibrary.utl.title"), plusWords)
	} else {
		plusAuthor := strings.Join(strings.Fields(author), "+")
		url = fmt.Sprintf(net.Koanf.String("openlibrary.utl.titleauthor"), plusWords, plusAuthor)
	}

	data, err := net.HTTPGet(url)
	if err != nil {
		return Response{}, err
	}

	var response Response
	json.Unmarshal(data, &response)

	for _, doc := range response.Docs {
		s, _ := json.MarshalIndent(doc, "", "\t")
		net.Logger.Trace(fmt.Sprintf("openlibrary: %s\n", string(s)))
	}

	return response, nil

}
