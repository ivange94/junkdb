package router

import (
	"io"
	"net/http"
)

type Ctx struct {
	w http.ResponseWriter
	r *http.Request
}

func (c *Ctx) Param(key string) string {
	return c.r.PathValue(key)
}

func (c *Ctx) Query(key string) string {
	return c.r.URL.Query().Get(key)
}

func (c *Ctx) Body() ([]byte, error) {
	return io.ReadAll(c.r.Body)
}
