package cmd

import (
	"os"
	"log"
	"strconv"
	"errors"
	"github.com/cdroot01/go01/note"
)


func Run() {
	args := os.Args
	
	if len(args) == 1 {
		log.Fatal("Enter command")
	}

	switch args[1] {
	case "list":
		note.PrintNotes()
	case "detail":
		index, err := getIndex(&args)
		if err != nil {
			log.Fatal(err)
		}

		err = note.PrintNote(index)
		if err != nil {
			log.Fatal(err)
		}
	case "create":
		log.Println("on create")
	case "delete":
		index, err := getIndex(&args)
		if err != nil {
			log.Fatal(err)
		}

		err = note.Delete(index)
		if err != nil {
			log.Fatal(err)
		}

		note.PrintNotes()
	default:
		log.Println("unknow command")
	}
}

func getIndex(args *[]string) (int, error) {
	if len(*args) == 2 {
		return 0, errors.New("Enter index note")
	}

	index, err := strconv.Atoi((*args)[2])
	if err != nil {
		return 0, errors.New("Index must be number")
	}

	return index, nil
}
