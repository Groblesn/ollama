package format

import (
	"fmt"
	"math"
	"strconv"
)

const (
	Thousand = 1000
	Million  = Thousand * 1000
	Billion  = Million * 1000
)

func HumanNumber(b uint64) string {
	switch {
	case b >= Billion:
		return formatScaled(float64(b)/Billion, 1, "B")
	case b >= Million:
		return formatScaled(float64(b)/Million, 2, "M")
	case b >= Thousand:
		return fmt.Sprintf("%.0fK", float64(b)/Thousand)
	default:
		return strconv.FormatUint(b, 10)
	}
}

// formatScaled drops the fractional part for whole numbers and otherwise
// keeps the given number of decimals.
func formatScaled(value float64, decimals int, suffix string) string {
	if value == math.Floor(value) {
		decimals = 0
	}
	return fmt.Sprintf("%.*f%s", decimals, value, suffix)
}
