package main

import (
	"fmt"
)

type Cache interface {
	Get(k string) (string, bool)
	Set(k, v string)
}

var _ Cache = (*cacheImpl)(nil)

// Доработает конструктор и методы кеша, так чтобы они соответствовали интерфейсу Cache
func newCacheImpl() *cacheImpl {
	return &cacheImpl{dbs: map[string]string{}}
}

type cacheImpl struct {
	dbs map[string]string
}

func (c *cacheImpl) Get(k string) (string, bool) {
	v, ok := c.dbs[k]
	return v, ok
}

func (c *cacheImpl) Set(k, v string) {
	c.dbs[k] = v
}

func newDbImpl(cache Cache) *dbImpl {
	return &dbImpl{cache: cache, dbs: map[string]string{"hello": "world", "test": "test"}}
}

type dbImpl struct {
	cache Cache
	dbs   map[string]string
}

func (d *dbImpl) Get(k string) (string, bool) {
	v, ok := d.cache.Get(k)
	if ok {
		return fmt.Sprintf("answer from cache: key: %s, val: %s", k, v), ok
	}

	v, ok = d.dbs[k]
	return fmt.Sprintf("answer from dbs: key: %s, val: %s", k, v), ok
}

func main() {
	c := newCacheImpl()
	c.Set("key1", "val1")
	c.Set("key2", "val2")
	c.Set("key3", "val3")
	db := newDbImpl(c)
	fmt.Println(db.Get("key1"))
	fmt.Println(db.Get("key2"))
	fmt.Println(db.Get("key3"))
	fmt.Println(db.Get("test"))
	fmt.Println(db.Get("hello"))
}
