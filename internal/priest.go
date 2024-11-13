package internal

import (
	"github.com/gone-io/gone"
	"github.com/gone-io/gone/goner"
)

//go:generate gonectr generate -m=../cmd/server
func Priest(cemetery gone.Cemetery) error {
	return goner.GinPriest(cemetery)
}
