package main

import (
	"fmt"
	"os"

	textprotogen "github.com/bookweb/textproto-gen"
)

func main() {
	if err := textprotogen.Run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}
}
