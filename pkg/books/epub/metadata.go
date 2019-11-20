package epub

import (
	"encoding/xml"
)

// Metadata --
type Metadata struct {
	XMLName    xml.Name `xml:"metadata"`
	Title      string   `xml:"title"`
	Language   string   `xml:"language"`
	Identifier []struct {
		Name   string `xml:"id,attr"`
		Scheme string `xml:"scheme,attr"`
		ID     string `xml:",innerxml"`
	} `xml:"identifier"`
	Creator     string `xml:"creator"`
	Contributor string `xml:"contributor"`
	Publisher   string `xml:"publisher"`
	Subject     string `xml:"subject"`
	Description string `xml:"description"`
	Event       []struct {
		Name string `xml:"event,attr"`
		Date string `xml:",innerxml"`
	} `xml:"date"`
	Type     string `xml:"type"`
	Format   string `xml:"format"`
	Source   string `xml:"source"`
	Relation string `xml:"relation"`
	Coverage string `xml:"coverage"`
	Rights   string `xml:"rights"`
}
