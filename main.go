package main

import (
	"log"

	"github.com/cdroot01/go01/cmd"
)

func init() {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
}

func main() {
	cmd.Run()
}
