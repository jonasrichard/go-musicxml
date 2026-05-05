package gomusicxml

import (
	"encoding/xml"
	"io"
)

// Read reads a ScorePartwise from the given reader.
// Currently only supports ScorePartwise and not TimePartwise.
func Read(r io.Reader) (*ScorePartwise, error) {
	var score ScorePartwise
	decoder := xml.NewDecoder(r)

	if err := decoder.Decode(&score); err != nil {
		return nil, err
	}

	return &score, nil
}

// Write writes the ScorePartwise to the given writer as XML.
func (s *ScorePartwise) Write(w io.Writer) error {
	encoder := xml.NewEncoder(w)
	encoder.Indent("", "  ")

	if err := encoder.Encode(s); err != nil {
		return err
	}

	return encoder.Flush()
}
