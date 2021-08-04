package http

import (
	"net/http/pprof"

	"github.com/gin-gonic/gin"
)

// Routes for pprof
func pprofRoutes(router *gin.Engine) {
	router.GET("/debug/pprof/", indexHandler())
	router.GET("/debug/pprof/heap", heapHandler())
	router.GET("/debug/pprof/goroutine", goroutineHandler())
	router.GET("/debug/pprof/block", blockHandler())
	router.GET("/debug/pprof/threadcreate", threadCreateHandler())
	router.GET("/debug/pprof/cmdline", cmdlineHandler())
	router.GET("/debug/pprof/profile", profileHandler())
	router.GET("/debug/pprof/symbol", symbolHandler())
	router.POST("/debug/pprof/symbol", symbolHandler())
}

func indexHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Index(ctx.Writer, ctx.Request)
	}
}

func heapHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Handler("heap").ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func goroutineHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Handler("goroutine").ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func blockHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Handler("block").ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func threadCreateHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Handler("threadcreate").ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func cmdlineHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Cmdline(ctx.Writer, ctx.Request)
	}
}

func profileHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Profile(ctx.Writer, ctx.Request)
	}
}

func symbolHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pprof.Symbol(ctx.Writer, ctx.Request)
	}
}
