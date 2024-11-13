package router

import "github.com/gone-io/gone"

const IdRouterPub = "router-pub"

//go:gone
func NewPubRouter() (gone.Goner, gone.GonerId) {
	return &pubRouter{}, IdRouterPub
}

type pubRouter struct {
	gone.Flag
	gone.IRouter
	root gone.RouteGroup `gone:"*"`
}

func (r *pubRouter) GetGonerId() gone.GonerId {
	return IdRouterPub
}

func (r *pubRouter) AfterRevive() gone.AfterReviveError {
	r.IRouter = r.root.Group("/api")
	return nil
}
