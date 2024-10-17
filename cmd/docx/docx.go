package docx

import (
	"encoding/xml"
	"io"
)

type R struct {
	T     *string `xml:"t"`
	DocPr *docPr  `xml:"drawing>inline>docPr"`
}
type docPr struct {
	Name *string `xml:"name,attr"`
}

func NewWRs(file io.Reader) []R {

	dec := xml.NewDecoder(file)
	wrs := make([]R, 0)
	for t, _ := dec.Token(); t != nil; t, _ = dec.Token() {
		switch doc := t.(type) {
		case xml.StartElement:
			var r R

			if doc.Name.Local == "r" {

				dec.DecodeElement(&r, &doc)
				wrs = append(wrs, r)
			}

		}
	}
	return wrs
}
func NewByteWRs(file io.Reader) []byte {
	wrs := NewWRs(file)
	byteWRs := make([]byte, 0)
	for _, r := range wrs {
		if r.T != nil {
			byteWRs = append(byteWRs, []byte(*r.T)...)
			byteWRs = append(byteWRs, []byte("\n")...)
		}
		if r.DocPr != nil {
			byteWRs = append(byteWRs, []byte("%Picture%")...)

			byteWRs = append(byteWRs, []byte(*r.DocPr.Name)...)

			byteWRs = append(byteWRs, []byte("\n")...)
		}

	}
	return byteWRs
}
