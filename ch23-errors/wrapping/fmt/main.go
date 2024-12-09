package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	err := transferFileContents("/my/imaginary/file")
	if err != nil {
		log.Printf("error occured: %s", err)
	}
}

func transferFileContents(filename string) error {
	contents, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("during file transfer impossible to open source file: %w", err)
	}
	err = os.WriteFile("/tmp/filecontents", contents, 0644)
	if err != nil {
		return fmt.Errorf("during file transfer impossible to write source file: %w", err)
	}
	return nil
}
