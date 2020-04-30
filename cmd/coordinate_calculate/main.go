package main

import (
	"fmt"

	"github.com/SinorStark/coordinateCalculate/internal/earth"
)

func main() {
	coordinate := earth.NewCoordinate(2020, 4, 30, 1.0, 2.0)
	fmt.Println(coordinate)
}
