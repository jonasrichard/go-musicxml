package gomusicxml

import (
	"encoding/xml"
	"fmt"
)

type Note struct {
	XMLName  xml.Name `xml:"note"`
	Pitch    Pitch    `xml:"pitch"`
	Duration int      `xml:"duration"`
	Type     string   `xml:"type"`
	// TODO model it correctly
	Rest      *struct{} `xml:"rest,omitempty"`
	Notations Notations `xml:"notations"`
}

type Pitch struct {
	Step   string `xml:"step"`
	Octave int    `xml:"octave"`
}

type Notations struct {
	Items []Notation `xml:",any"`
}

type Notation interface {
	isNotation()
}

type Technical struct {
	XMLName  xml.Name  `xml:"technical"`
	Fret     *int      `xml:"fret,omitempty"`
	HammerOn *HammerOn `xml:"hammer-on,omitempty"`
	String   *int      `xml:"string,omitempty"`
	Extra    string    `xml:",innerxml"`
}

type HammerOn struct {
	Number int    `xml:"number,attr"`
	Text   string `xml:",chardata"`
	Type   string `xml:"type,attr"`
}

func (Technical) isNotation() {}

func (ns *Notations) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	fmt.Printf("marshalling notations: start element: %v\n", start)

	if err := e.EncodeToken(start); err != nil {
		return err
	}

	for _, n := range ns.Items {
		switch v := n.(type) {
		case Technical:
			if err := e.Encode(v); err != nil {
				return err
			}
		default:
			return nil
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (ns *Notations) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch t := tok.(type) {
		case xml.StartElement:
			var n Notation
			switch t.Name.Local {
			case "technical":
				var tech Technical
				if err := d.DecodeElement(&tech, &t); err != nil {
					return err
				}
				n = tech
			}
			ns.Items = append(ns.Items, n)
		case xml.EndElement:
			if t.Name.Local == start.Name.Local {
				return nil
			}
		}
	}
}
