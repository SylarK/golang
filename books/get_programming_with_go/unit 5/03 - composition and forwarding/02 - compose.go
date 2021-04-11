package main

import "fmt"

type report struct {
	sol         int
	temperature temperature
	location    location
}

type temperature struct {
	high, low celsius
}

type location struct {
	lat, long float64
}

type celsius float64

func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}

func main() {

	bradbury := location{-4.8596, 137.4469}
	t := temperature{high: -1.0, low: -78.0}
	report := report{sol: 15, temperature: t, location: bradbury}

	fmt.Printf("%v\n", report)

	fmt.Printf("a balmy %v C\n", report.temperature.high)
	fmt.Printf("average %v C\n", report.temperature.average())
	fmt.Println(t.average())
}
