package rectangle

import "math"
import "fmt"

func init() {
	fmt.Println("rectangle package initialized")
}

// Area of rectangle with given length and width
func Area(len, wid float64) float64 {
	area := len * wid
	return area
}

// Diagonal length of rectangle with given length and width
func Diagonal(len, wid float64) float64 {
	diagonal := math.Sqrt((len * len) + (wid * wid))
	return diagonal
}
