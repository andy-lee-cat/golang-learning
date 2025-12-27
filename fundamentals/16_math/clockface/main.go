package main

import (
	"os"
	"time"

	"golang-learning/fundamentals/16_math/svg"
)

func main() {
	t := time.Now()
	svg.Write(os.Stdout, t)
}
