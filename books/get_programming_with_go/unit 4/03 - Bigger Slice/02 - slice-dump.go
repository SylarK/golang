package main

import "fmt"

func dump(label string, slice []string) {
	fmt.Printf("%v: length %v, capacity %v %v\n", label, len(slice), cap(slice), slice)
}

func main() {
	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dwarfs1 := append(dwarfs, "Orcus")
	dwarfs2 := append(dwarfs1, "Salacia", "Quaoar", "Sedna", "Pan")
	dump("dwarfs", dwarfs)
	dump("dwarfs[1:2]", dwarfs[1:2])
	dump("dwarfs1", dwarfs1)
	dump("dwarfs2", dwarfs2)
}
