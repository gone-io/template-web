package main

import (
	"github.com/gone-io/gone"
	"tempalte_model/internal"
)

func main() {
	gone.
		Default.
		LoadPriest(internal.Priest).
		Serve()
}
