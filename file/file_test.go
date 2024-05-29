package file

import "testing"

func TestOpen(t *testing.T) {
	lines, err := Open("./data.tx")

	if err != nil {
		t.Fatal(err)
	}

	t.Log(lines)
}

func TestWrite(t *testing.T) {
	lines := []string{
		"Vip pro",
		"Go to new line",
	}

	err := Write(lines, "./data.txt")

	if err != nil {
		t.Fatal(err)
	}
}