package main

import (
	"fmt"
	"math"

	strutil "github.com/odobr/go_crash_course/03_packages/str"
)

func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(9))
	fmt.Println(strutil.Reverse("hello"))
}
