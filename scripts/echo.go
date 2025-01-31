package main

import (
	"fmt"
	"io"

	"github.com/bitfield/script"
)

func main() {
	script.Echo("a\nb\nc").FilterScan(func(line string, w io.Writer) {
		fmt.Fprintf(w, "scanned line: %q\n", line)
	}).Stdout()

	f, err := script.File("go.mod").String()
	if err != nil {
		panic(err)
	}
	fmt.Println(f)
}
