package caloryman

import (
	"fmt"
	"strconv"
	"strings"
)

func UCFrom(unit string) float64 {
	x := strings.Split(unit, " ")
	v, _ := strconv.ParseFloat(x[0], 64)

	if len(x) <= 1 {
		return v
	}
	if strings.ToLower(x[0]) == "mg" {
		return v * 0.001
	}
	if strings.ToLower(x[0]) == "ug" {
		return v * 0.000001
	}

	return v
}

func UCTo(unit float64, mode string) string {
	return fmt.Sprintf("%.1f g", unit)
}
