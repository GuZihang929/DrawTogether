package api

import (
	"context"
	"draw-together/server/api"
)

func Start(c context.Context) {
	Router := api.Routers()
	s := api.InitServer(Router)
	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
	api.Shutdown(s)
}
