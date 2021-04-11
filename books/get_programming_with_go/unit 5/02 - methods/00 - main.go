package main

import "fmt"

type coordinate struct {
	//degrees, minutes, seconds in a N/S/E/W hemisphere.
	d, m, s float64
	h       rune
}

type location struct {
	lat, long float64
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}

	return sign * (c.d + c.m/60 + c.s/3600)
}

func main() {

	lat := coordinate{4, 35, 22.2, 'S'}
	long := coordinate{137, 26, 30.12, 'E'}

	curiosity := location{lat.decimal(), long.decimal()}

	fmt.Println(lat.decimal(), long.decimal())
	fmt.Println(curiosity)
}
