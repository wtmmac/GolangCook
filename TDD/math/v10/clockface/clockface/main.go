package main

import (
	"os"
	"time"

	"github.com/wtmmac/GolangCook/TDD/math/v10/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
