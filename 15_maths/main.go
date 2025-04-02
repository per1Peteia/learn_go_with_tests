package main

import (
	"github.com/per1Peteia/learn_go_with_tests/15_maths/clockface"
	"os"
	"time"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
