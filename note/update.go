package note

func Delete(index int) error {
	err := ValidIndex(index)

	if err != nil {
		return err
	}

	Notes = append(Notes[:index], Notes[index:]...)
	return nil
}
