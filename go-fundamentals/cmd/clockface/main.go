package main

import (
	"log"
	"os"
	"time"

	"github.com/cativovo/learn-go-with-tests/math/clockface"
)

func main() {
	t := time.Now()
	if err := clockface.SVGWriter(os.Stdout, t); err != nil {
		log.Fatal(err)
	}
}
