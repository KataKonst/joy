package cachestorage

import (
	"github.com/matthewmueller/golly/dom/cache"
	"github.com/matthewmueller/golly/dom/cachequeryoptions"
	"github.com/matthewmueller/golly/dom/request"
	"github.com/matthewmueller/golly/js"
)

// CacheStorage struct
// js:"CacheStorage,omit"
type CacheStorage struct {
}

// Delete fn
// js:"delete"
func (*CacheStorage) Delete(cacheName string) (b bool) {
	js.Rewrite("await $_.delete($1)", cacheName)
	return b
}

// Has fn
// js:"has"
func (*CacheStorage) Has(cacheName string) (b bool) {
	js.Rewrite("await $_.has($1)", cacheName)
	return b
}

// Keys fn
// js:"keys"
func (*CacheStorage) Keys() (s []string) {
	js.Rewrite("await $_.keys()")
	return s
}

// Match fn
// js:"match"
func (*CacheStorage) Match(request *request.Request, options *cachequeryoptions.CacheQueryOptions) (i interface{}) {
	js.Rewrite("await $_.match($1, $2)", request, options)
	return i
}

// Open fn
// js:"open"
func (*CacheStorage) Open(cacheName string) (c *cache.Cache) {
	js.Rewrite("await $_.open($1)", cacheName)
	return c
}
