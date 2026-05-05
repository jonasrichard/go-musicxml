package gomusicxml

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestNoteUnmarshalTechnicalProcInst(t *testing.T) {
	input := `
	<note>
		<pitch><step>A</step><octave>4</octave></pitch>
		<duration>1</duration>
		<type>quarter</type>
		<notations>
			<technical><?GP <root><string>6</string><fret>12</fret></root>?></technical>
		</notations>
	</note>`

	var note Note
	if err := xml.Unmarshal([]byte(input), &note); err != nil {
		t.Fatalf("unmarshal note: %v", err)
	}

	if len(note.Notations.Items) != 1 {
		t.Fatalf("expected 1 notation, got %d", len(note.Notations.Items))
	}

	technical, ok := note.Notations.Items[0].(Technical)
	if !ok {
		t.Fatalf("expected Technical notation, got %T", note.Notations.Items[0])
	}

	if got := technical.Extra; got != `<?GP <root><string>6</string><fret>12</fret></root>?>` {
		t.Fatalf("expected technical extra data, got %q", got)
	}
}

func TestNoteMarshalTechnicalProcInst(t *testing.T) {
	note := Note{
		Pitch:    Pitch{Step: "A", Octave: 4},
		Duration: 1,
		Type:     "quarter",
		Notations: Notations{
			Items: []Notation{
				Technical{Extra: `<?GP <root><string>6</string><fret>12</fret></root> ?>`},
			},
		},
	}

	buf, err := xml.Marshal(note)
	if err != nil {
		t.Fatalf("marshal note: %v", err)
	}

	out := string(buf)
	if !strings.Contains(out, `<notations><technical><?GP <root><string>6</string><fret>12</fret></root> ?></technical></notations>`) {
		t.Fatalf("expected technical PI in notations, got %s", out)
	}
}

func TestUnmarshalHammerOn(t *testing.T) {
	input := `
	<note>
		<pitch><step>A</step><octave>4</octave></pitch>
		<duration>1</duration>
		<type>quarter</type>
		<notations>
			<technical><hammer-on number="1" type="start">H</hammer-on></technical>
		</notations>
	</note>`

	var note Note
	if err := xml.Unmarshal([]byte(input), &note); err != nil {
		t.Fatalf("unmarshal note: %v", err)
	}

	if len(note.Notations.Items) != 1 {
		t.Fatalf("expected 1 notation, got %d", len(note.Notations.Items))
	}

	technical, ok := note.Notations.Items[0].(Technical)
	if !ok {
		t.Fatalf("expected Technical notation, got %T", note.Notations.Items[0])
	}

	if technical.HammerOn == nil {
		t.Fatal("expected HammerOn notation, got nil")
	}

	if got := technical.HammerOn.Number; got != 1 {
		t.Fatalf("expected hammer-on number 1, got %d", got)
	}

	if got := technical.HammerOn.Type; got != "start" {
		t.Fatalf("expected hammer-on type start, got %q", got)
	}

	if got := technical.HammerOn.Text; got != "H" {
		t.Fatalf("expected hammer-on text H, got %q", got)
	}
}
