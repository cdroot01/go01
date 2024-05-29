package note

import (
	"testing"
)

func TestGetNotes(t *testing.T) {
	PrintNotes()
}

func TestDetailNote(t *testing.T) {
	note, err := Detail(10)

	if err != nil {
		t.Fatal(err)
	}
	
	t.Log(note)
}

func TestOpen(t *testing.T) {
	notes, err := Open()
	
	if err != nil {
		t.Fatal(err)
	}

	t.Log(notes)
}