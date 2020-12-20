package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	//pepe()
	//fmt.Print(conversion)
	pow(3, 2, 5)
}

func pepe() {
	fmt.Print("Hola soy pepe")
}

func conversion() (int, error) {

	palabra := "dos"
	conversion, error := strconv.Atoi(palabra)

	if error != nil {
		return 0, error
	}
	return conversion, nil
	//numero := 1
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("miren esta cuentita: %g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

// ahi me fije, en python es igual: str(...), int(...)
