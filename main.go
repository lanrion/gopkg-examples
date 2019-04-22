package main

import (
	"os"

	"github.com/lanrion/gopkg-examples/syncT"
)

func main() {
	syncT.Log(os.Stdout, "path", "/search?q=flowers")
}
