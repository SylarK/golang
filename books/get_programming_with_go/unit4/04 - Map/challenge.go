/*

Write a function to count the frequency of words in a string of text and return a map
of words with their counts. The function should convert the text to lowercase, and
punctuation should be trimmed from words. The strings package contains several
helpful functions for this task, including Fields, ToLower, and Trim.

*/

package main

import (
	"fmt"
	"strings"
)

func main() {

	phrase := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed cursus augue mi, a rutrum urna ultricies ut. " +
		" Proin ornare justo ac dolor interdum hendrerit. "
	//phrase := "As far as eye could reach he saw nothing but the stems of the great plants about him	receding in the violet shade, and far overhead the multiple transparency of huge leaves 	filtering the sunshine to the solemn splendour of twilight in which he walked. Whenever	he felt able he ran again; the ground continued soft and springy, covered with the same	resilient weed which was the first thing his hands had touched in Malacandra. Once or	twice a small red creature scuttled across his path, but otherwise there seemed to be no	life stirring in the wood; nothing to fear—except the fact of wandering unprovisioned	and alone in a forest of unknown vegetation thousands or millions of miles beyond the	reach or knowledge of man."
	phrase = strings.ToLower(phrase)
	phrase = strings.ReplaceAll(phrase, " ", "")

	col := strings.Split(phrase, "")

	frequency := make(map[string]int)

	for _, v := range col {
		frequency[v]++
	}

	for i, v := range frequency {

		fmt.Printf("%v\t:\t%v\n", i, v)
	}
}
