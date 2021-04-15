package main

import (
	"fmt"
	"os"
)

func proverbs(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}

	// defer f.Close()

	_, err = fmt.Fprintln(f, "Error are values.")
	if err != nil {
		f.Close()
		return err
	}

	_, err = fmt.Fprintln(f, "Don't just check errors, handle them gracefully.")
	f.Close()
	return err
}

func main() { fmt.Println() }
