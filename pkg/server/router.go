package server

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/NganJason/BE-template/pkg/clog"
	"github.com/gorilla/mux"
)

type Route struct {
	Method  string
	Path    string
	Name    string
	Handler Handler
	Req     interface{}
}

type Router struct {
	muxR           *mux.Router
	middlewareList []Middleware
}

var mainRouter *Router

func NewMainRouter() *Router {
	if mainRouter != nil {
		return mainRouter
	}

	mainRouter = &Router{
		muxR: mux.NewRouter(),
	}

	return mainRouter
}

func (er *Router) AddRoute(route *Route) {
	er.muxR.HandleFunc(route.Path, er.toHttpHandler(route.Handler, route))
}

func (er *Router) AddMiddleware(m Middleware) {
	er.middlewareList = append(er.middlewareList, m)
}

func (er *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	er.muxR.ServeHTTP(w, r)
}

func (er *Router) toHttpHandler(handler Handler, route *Route) http.HandlerFunc {
	respHandler := func(ctx context.Context, writer http.ResponseWriter, req *http.Request, resp *HandlerResp) *HandlerResp {
		return resp
	}

	for idx := len(er.middlewareList) - 1; idx >= 0; idx-- {
		m := er.middlewareList[idx]
		if m.CanSkip(route.Name) {
			continue
		}

		handler = m.PreRequest(handler)
		respHandler = m.PostRequest(respHandler)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := clog.ContextWithTraceID(
			r.Context(),
			strconv.FormatUint(uint64(time.Now().Unix()), 10),
		)

		ctx = context.WithValue(ctx, CtxRequestStruct, route.Req)

		tmBegin := time.Now()
		clog.Infof(
			ctx,
			"request started|traceID=%s | url=%s",
			clog.GetTraceID(ctx), r.URL.Path,
		)

		resp := handler(ctx, w, r)
		if resp != nil {
			resp = respHandler(ctx, w, r, resp)

			tmEnd := time.Now()
			clog.Infof(
				ctx,
				"request ended|resp=%+v | err=%v | proctm=%dµs",
				resp.Payload, resp.Err, tmEnd.Sub(tmBegin).Nanoseconds()/1000,
			)

			JsonResponse(w, resp.Payload, resp.Err)
		}
	}
}
