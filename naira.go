package naira

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

/*
Inspired by this StackOverflow answer from Niko Kovacevic
https://stackoverflow.com/a/45472066
*/

// Kobo represents a naira amount in kobo
type Kobo int64

//Tokobo Converts a float based naira amount to kobo, second argument should be "d", "u", or "n"
// to round result down, up, or to nearest kobo
func ToKobo(f float64, direction string) Kobo {
	x := rounder(direction, f)
	return Kobo(x * 100)
}

// KoboToFloat converts a Kobo amount to float64 based naira
func (m Kobo) KoboToFloat() float64 {
	x := float64(m)
	x = x / 100
	return x
}

// Multiply safely multiplies a Kobo value by a float64
func (m Kobo) Multiply(f float64) Kobo {
	x := float64(m) * f
	return Kobo(x)
}

// Divide safely divides a Kobo value by a float64
func (m Kobo) Divide(f float64) Kobo {
	x := float64(m) / f
	return Kobo(x)
}

// KobotoPrettyNGNString returns a pretty NGN value
func (m Kobo) KobotoPrettyNGNString() string {

	kstr := strconv.FormatInt(int64(m), 10) //1. Convert to string

	ksl := strings.Split(kstr, "") //2. Split into []string

	lksl := len(ksl) - 1
	decimals := []string{".", ksl[lksl-1], ksl[lksl]} // handle decimals

	var outputslice []string
	for i, v := lksl-2, 1; v < lksl; i-- {
		// prepend current value of ksl to output slice
		outputslice = append([]string{ksl[i]}, outputslice...)
		// if position of lksl % 3 == 0, also prepend a comma
		if v%3 == 0 {
			if i == lksl-2 || i == 0 || // Except when it's the last digit before the decimals or first digit
				ksl[i-1] == "-" { // ...or when the first digit is a minus sign
				v++
				continue
			}
			outputslice = append([]string{","}, outputslice...)
		}
		v++
	}
	outputslice = append(outputslice, decimals...)
	//6. Join
	return strings.Join(outputslice, "")
}

//IntStrInNairaToKobo converts an integer string to kob
func IntStrInNairaToKobo(is string) (Kobo, error) {
	// convert string to int
	i, err := strconv.Atoi(is)
	if err != nil {
		return 0, errors.New("couldn't convert to kobo: " + err.Error())
	}
	// convert int to kobo
	return Kobo(i * 100), nil
}

func rounder(direction string, f float64) float64 {
	var x float64
	switch direction {
	case "d":
		x = math.Floor(f*100) / 100
	case "n":
		x = math.Round(f*100) / 100
	case "u":
		x = math.Ceil(f*100) / 100
	}
	return x
}
