package file

import (
	"bufio"
	"os"
	"strings"
)

func Open(path string) ([]string, error){
	var lines []string

	file, err := os.Open(path)
	
	if err != nil {
		return lines, err
	}
	defer file.Close()

	
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	
	return lines, nil
}

func Write(lines []string, path string) error {
    data := []byte(strings.Join(lines, "\n")) // Join lines with newlines
    err := os.WriteFile(path, data, 0644) // 0644 gives read/write permissions
	return err
}