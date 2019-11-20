package epub

import (
	"encoding/xml"
)

// Manifest element provides an exhaustive list of the Publication Resources that constitute
// the EPUB Publication, each represented by an item element.
type Manifest struct {
	XMLName xml.Name `xml:"manifest"`
	ID      string   `xml:"id,attr,omitempty"` // The ID [XML] of this element, which must be unique within the document scope.
	Items   []*struct {
		ID           string `xml:"id,attr"`                      // The ID [XML] of this element, which must be unique within the document scope.
		Href         string `xml:"href,attr"`                    // An IRI [RFC3987] specifying the location of the Publication Resource described by this item.
		MediaType    string `xml:"media-type,attr"`              // A media type [RFC2046] that specifies the type and format of the Publication Resource described by this item.
		Fallback     string `xml:"fallback,attr,omitempty"`      // An IDREF [XML] that identifies the fallback for a non-Core Media Type.
		Properties   string `xml:"properties,attr,omitempty"`    // A space-separated list of property values.
		MediaOverlay string `xml:"media-overlay,attr,omitempty"` // An IDREF [XML] that identifies the Media Overlay Document for the resource described by this item.
	} `xml:"item"` // List of the Publication Resources
}
