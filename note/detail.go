package note

import "fmt"

func Detail(index int) (string, error ) {
	err := ValidIndex(index)

	if err != nil {
		return "", err
	}

	return Notes[index], nil
}

func PrintNotes() {
	for i, v := range Notes {
		fmt.Printf("%v: %v\n", i, v)
	}
}

func PrintNote(index int) error {
	note, err := Detail(index)

	if err != nil {
		return err
	}

	fmt.Printf("%v: %v", index, note)
	
	return nil
}
