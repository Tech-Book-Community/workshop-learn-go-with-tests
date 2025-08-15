package main

import (
	"os"
	"time"

	"github.com/Tech-Book-Community/workshop-learn-go-with-tests/13-math-acceptance-test/gson/svg"
)

func main() {
	t := time.Now()
	svg.Write(os.Stdout, t)
}
