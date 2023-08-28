package cache

import "errors"

type Cache struct {
	cache map[string]any
}

func New() *Cache {
	c := make(map[string]any)
	return &Cache{c}
}

func (c *Cache) Set(key string, val any) error {
	_, exists := c.cache[key]
	if exists {
		return errors.New("A value with this key already exists [" + key + "]")
	}
	c.cache[key] = val
	return nil
}

func (c *Cache) Get(key string) (any, error) {
	val, exists := c.cache[key]
	if !exists {
		return nil, errors.New("No value with this key [" + key + "]")
	}
	return val, nil
}

func (c *Cache) Delete(key string) error {
	_, exists := c.cache[key]
	if !exists {
		return errors.New("No value with this key [" + key + "]")
	}
	delete(c.cache, key)
	return nil
}
