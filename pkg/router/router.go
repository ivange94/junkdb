package router

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type Router struct {
	mux          *http.ServeMux
	errorHandler ErrorHandler
	timeout      time.Duration
}

type ErrorHandler func(*Ctx, error)

type HandlerFunc func(*Ctx) error

type Option func(*Router)

func New(opts ...Option) *Router {
	r := &Router{}
	for _, opt := range opts {
		opt(r)
	}
	if r.mux == nil {
		r.mux = http.NewServeMux()
	}
	if r.errorHandler == nil {
		r.errorHandler = defaultErrorHandler
	}
	return r
}

func WithServeMux(mux *http.ServeMux) Option {
	return func(r *Router) {
		r.mux = mux
	}
}

func WithErrorHandler(he ErrorHandler) Option {
	return func(r *Router) {
		r.errorHandler = he
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(r *Router) {
		r.timeout = timeout
	}
}

func defaultErrorHandler(*Ctx, error) {}

func (r *Router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	r.mux.ServeHTTP(writer, request)
}

func (r *Router) handle(method, path string, h HandlerFunc) {
	pattern := fmt.Sprintf("%s %s", method, path)

	r.mux.HandleFunc(pattern, func(w http.ResponseWriter, req *http.Request) {
		if r.timeout > 0 {
			ctx, cancel := context.WithTimeout(req.Context(), r.timeout)
			defer cancel()

			req = req.WithContext(ctx)
		}
		ctx := &Ctx{w: w, r: req}
		if err := h(ctx); err != nil {
			r.errorHandler(ctx, err)
		}
	})
}

func (r *Router) Get(path string, handler HandlerFunc) {
	r.handle(http.MethodGet, path, handler)
}

func (r *Router) Post(path string, handler HandlerFunc) {
	r.handle(http.MethodPost, path, handler)
}

func (r *Router) Put(path string, handler HandlerFunc) {
	r.handle(http.MethodPut, path, handler)
}

func (r *Router) Delete(path string, handler HandlerFunc) {
	r.handle(http.MethodPost, path, handler)
}

func (r *Router) Patch(path string, handler HandlerFunc) {
	r.handle(http.MethodPatch, path, handler)
}
