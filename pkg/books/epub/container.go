package epub

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"
)

// Default file names and mime-type.
var (
	RootPath        = "OEBPS"                // Folder with content of publication
	PackageFilename = "package.opf"          // Package description file name
	EPUBMimeType    = "application/epub+zip" // EPUB mime-type
)

// Predefined names of folder and container file.
const (
	METAINF      = "META-INF"      // Predefined net folder
	CONTAINER    = "container.xml" // Predefined container file name
	containerXML = METAINF + "/" + CONTAINER
)

// Container describes the contents of the container
type Container struct {
	XMLName xml.Name `xml:"urn:oasis:names:tc:opendocument:xmlns:container container"`
	Version string   `xml:"version,attr"`
	Package
	Rootfiles []RootFile `xml:"rootfiles>rootfile"`
}

// RootFile describes the path to description of publication.
type RootFile struct {
	FullPath  string `xml:"full-path,attr"`
	MediaType string `xml:"media-type,attr"`
}

// New -- read a file and return a container
func New(filename string) (Container, error) {
	ebook := Container{}

	zipfile, err := os.Open(filename)
	if err != nil {
		return ebook, fmt.Errorf("Cannot open %q: %w", filename, err)
	}
	defer zipfile.Close()

	zipfileStat, err := zipfile.Stat()
	if err != nil {
		return ebook, fmt.Errorf("Cannot stat %q: %w", filename, err)
	}

	zipReader, err := zip.NewReader(zipfile, zipfileStat.Size())
	if err != nil {
		return ebook, fmt.Errorf("Cannot parse zipfile %q: %w", filename, err)
	}

	files := make(map[string]*zip.File)
	for _, f := range zipReader.File {
		files[f.Name] = f
	}

	container, err := files[containerXML].Open()
	if err != nil {
		return ebook, fmt.Errorf("Cannot open %s: %s", containerXML, err.Error())
	}

	var buffer bytes.Buffer
	_, err = io.Copy(&buffer, container)
	if err != nil {
		return ebook, fmt.Errorf("Cannot read container: %s", err.Error())
	}

	err = xml.Unmarshal(buffer.Bytes(), &ebook)
	if err != nil {
		return ebook, fmt.Errorf("Cannot unmarshal container: %s", err.Error())
	}

	if len(ebook.Rootfiles) < 1 {
		return ebook, errors.New("ebook: no rootfile found in container")
	}

	for _, rootfile := range ebook.Rootfiles {
		if files[rootfile.FullPath] == nil {
			return ebook, errors.New("No rootfile " + rootfile.FullPath)
		}

		rootFile, err := files[rootfile.FullPath].Open()
		if err != nil {
			return ebook, fmt.Errorf("Cannot open rootfile %s: %s", rootfile.FullPath, err.Error())
		}

		var buffer bytes.Buffer
		_, err = io.Copy(&buffer, rootFile)
		if err != nil {
			return ebook, fmt.Errorf("Cannot read rootfile %s: %s", rootfile.FullPath, err.Error())
		}

		err = xml.Unmarshal(buffer.Bytes(), &ebook.Package)
		if err != nil {
			return ebook, fmt.Errorf("Cannot parse rootfile %s: %s", rootfile.FullPath, err.Error())
		}
	}

	return ebook, nil
}
