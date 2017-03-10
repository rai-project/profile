package web

import (
	"net/http"
	"net/http/pprof"

	"github.com/labstack/echo"
)

func Init(api *echo.Group) {
	wrap := func(f func(w http.ResponseWriter, r *http.Request)) echo.HandlerFunc {
		return echo.WrapHandler(http.HandlerFunc(f))
	}
	api.GET("/debug/pprof/", wrap(pprof.Index))
	api.GET("/debug/pprof/cmdline", wrap(pprof.Cmdline))
	api.GET("/debug/pprof/profile", wrap(pprof.Profile))
	api.GET("/debug/pprof/symbol", wrap(pprof.Symbol))

	// Manually add support for paths linked to by index page at /debug/pprof/
	api.GET("/debug/pprof/goroutine", echo.WrapHandler(pprof.Handler("goroutine")))
	api.GET("/debug/pprof/heap", echo.WrapHandler(pprof.Handler("heap")))
	api.GET("/debug/pprof/threadcreate", echo.WrapHandler(pprof.Handler("threadcreate")))
	api.GET("/debug/pprof/block", echo.WrapHandler(pprof.Handler("block")))
}
