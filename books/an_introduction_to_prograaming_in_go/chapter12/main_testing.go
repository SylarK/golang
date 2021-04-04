/*
	Go includes a special program that makes writing tests easier.
	"go test" command will look for any tests in any of the files in 
	the current folder and run them.

	Tests are indetified by starting a function with the word Test and
	taking one argument of type "*testing.T".
*/

package math

import "testing"

type testpair struct{
	values []float64
	average float64
}

var tests = []testpair{
	{ []float64{1,2},1.5},
	{ []float64{1,,1,1,1,1,1,1},1},
	{ []float64{-1,1},0},
}

funct TestAverage(t *testing.T){
	for _, pair := range tests{
		v := Average(pair.values)
		if v != pair.average{
			t.Error(
				"For", pair.values,
				"expected", pair.average,
				"got", v,
			)
		}
	}
}