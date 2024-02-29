package main

import (
	"os"
	"time"

	"github.com/wtmmac/GolangCook/TDD/math/v11/clockface"
)

func main() {
	t := time.Now()
	// fmt.Println(t)
	clockface.SVGWriter(os.Stdout, t)
}
