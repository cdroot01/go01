package main

import (
	"testing"
	"bufio"
	"os"
)

func TestSomething(t *testing.T){
	t.Log("on testing")
}

func TestOpenFile(t *testing.T) {
	file, err := os.Open("./data.txt")
	
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	
	t.Log(lines)
}
