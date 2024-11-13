package main

import (
	"web/internal"

	"github.com/gone-io/gone"
)

func main() {
	gone.
		Default.
		LoadPriest(internal.Priest).
		Serve()
}
