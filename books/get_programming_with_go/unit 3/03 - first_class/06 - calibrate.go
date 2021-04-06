package main

import "fmt"

type kelvin float64
type sensor func() kelvin

func realSensor() kelvin {
	return 0
}

func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}

func main() {
	sensor := calibrate(realSensor, 5)
	fmt.Println(sensor())

	var k kelvin = 294.0

	sensor_fake := func() kelvin {
		return k
	}

	fmt.Println((sensor_fake()))
	k++
	fmt.Println((sensor_fake()))
}
