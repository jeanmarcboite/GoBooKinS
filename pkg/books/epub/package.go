package epub

import (
	"encoding/xml"
)

// Package element is the root container of the Package Document and encapsulates Publication
// metadata and resource information.
type Package struct {
	XMLName          xml.Name `xml:"http://www.idpf.org/2007/opf package"`
	Version          string   `xml:"version,attr"`            // Specifies the EPUB specification version to which the Publication conforms
	UniqueIdentifier string   `xml:"unique-identifier,attr"`  // An IDREF that identifies the dc:identifier element that provides the package's preferred, or primary, identifier
	Prefix           string   `xml:"prefix,attr,omitempty"`   // Declaration mechanism for prefixes not reserved by this specification.
	Lang             string   `xml:"xml:lang,attr,omitempty"` // Specifies the language used in the contents and attribute values of the carrying element and its descendants
	Dir              string   `xml:"dir,attr,omitempty"`      // Specifies the base text direction of the content and attribute values of the carrying element and its descendants.
	ID               string   `xml:"id,attr,omitempty"`       // The ID of this element, which must be unique within the document scope
	Metadata         Metadata `xml:"metadata"`
	Manifest         Manifest `xml:"manifest"`
	Spine            Spine    `xml:"spine"`
	Guide            Guide
}
// Guide --
type Guide struct {
	// <reference href="Text/content0001.xhtml" title="Cover Page" type="cover" />
}