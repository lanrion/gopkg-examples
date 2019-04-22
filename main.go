package main

import (
	"os"

	sync "github.com/lanrion/gopkg-examples/sync"
)

func main() {
	sync.Log(os.Stdout, "path", "/search?q=flowers")
}
