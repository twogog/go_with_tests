package main

import (
	"os"
	"time"

	"golang/math/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
