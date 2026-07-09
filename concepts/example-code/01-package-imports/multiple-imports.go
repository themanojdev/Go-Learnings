package packageimports
import (
	"fmt"
	"strings"
	"math"
)

func MultipleImports() {
	fmt.Println("=== Multiple Imports Example ===")

	fmt.Println("***Using Strings Package***")
	text := "golang"
	text = strings.ToUpper(text)
	fmt.Println("Text variable changed to upper case",text)

	fmt.Println("***Using Math Package***")
	var num float64 = 16
	var result = math.Sqrt(num)
	fmt.Println("Printing the square root of num variable",result)
}