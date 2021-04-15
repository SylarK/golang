package main

import "fmt"

func mirepoix(ingredients []string) []string {
	if ingredients == nil {
		return append(ingredients, "onion", "carrot", "celery")
	}
	return append(ingredients, "mac", "che", "bah", "tah")

}

func main() {

	soup := mirepoix(nil)
	for _, ingredient := range soup {
		fmt.Println(ingredient)
	}
}
