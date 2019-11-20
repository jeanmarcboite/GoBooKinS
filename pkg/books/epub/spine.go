package epub

import "encoding/xml"

// Spine element defines the default reading order of the EPUB Publication content by defining
// an ordered list of manifest item references.
type Spine struct {
	XMLName       xml.Name `xml:"spine"`
	ID            string   `xml:"id,attr,omitempty"`                         // The ID [XML] of this element, which must be unique within the document scope.
	Toc           string   `xml:"toc,attr,omitempty"`                        // An IDREF [XML] that identifies the manifest item that represents the superseded NCX.
	PageDirection string   `xml:"page-progression-direction,attr,omitempty"` // The global direction in which the Publication content flows. Allowed values are ltr (left-to-right), rtl (right-to-left) and default.
	ItemRefs      []*struct {
		IDRef      string `xml:"idref,attr"`                // An IDREF [XML] that identifies a manifest item.
		Linear     string `xml:"linear,attr,omitempty"`     // Specifies whether the referenced content is primary. The value of the attribute must be yes or no. The default value is yes.
		ID         string `xml:"id,attr,omitempty"`         // The ID [XML] of this element, which must be unique within the document scope.
		Properties string `xml:"properties,attr,omitempty"` // A space-separated list of property values.
	} `xml:"itemref"` // Ordered subset of the Publication Resources listed in the manifest
}
