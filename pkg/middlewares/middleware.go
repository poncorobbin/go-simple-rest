package middlewares

import "net/http"

type CustomMux struct {
	http.ServeMux
	middlewares []func(http.Handler) http.Handler
}

func (c *CustomMux) RegisterMiddleware(next func(http.Handler) http.Handler) {
	c.middlewares = append(c.middlewares, next)
}

func (c *CustomMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var handler http.Handler = &c.ServeMux

	for _, next := range c.middlewares {
		handler = next(handler)
	}

	handler.ServeHTTP(w, r)

}
