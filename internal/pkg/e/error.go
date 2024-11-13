package e

import (
	"net/http"

	"github.com/gone-io/gone"
)

var (
	ErrUnauthorized = gone.NewError(1001, "没有登录或登录失效", http.StatusUnauthorized)
)
