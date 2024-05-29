package note

import (
	"errors"
	"github.com/cdroot01/go01/file"
)

var Notes []string

func ValidIndex(index int) error {
	if index < 0 {
		return errors.New("Index must greater than 0")
	}
	if index > len(Notes) - 1 {
		return errors.New("Out of index")
	}

	return nil
}

func Open() ([]string , error) {
	return file.Open("./data.txt")
}