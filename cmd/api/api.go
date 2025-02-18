package api

import (
	"context"
	"draw-together/server/api"
)

func Start(c context.Context) {

	// 服务发现

	Router := api.Routers()
	s := api.InitServer(Router)
	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
	api.Shutdown(s)
}
