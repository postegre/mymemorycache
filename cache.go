package cache

type Cache struct {
	data map[string]interface{}
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]interface{}),
	}
}

func (c *Cache) Set(key string, value interface{}) {
	c.data[key] = value
}

func (c *Cache) Get(key string) interface{} {
	value := c.data[key]
	return value
}

func (c *Cache) Delete(key string) {
	delete(c.data, key)
}
