package cache

type Cache struct {
	data map[string]interface{}
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]interface{}),
	}
}

func Set(key string, value interface{})

func (c *Cache) Get(key string) interface{} {
	value := c.data[key]
	return value
}

func (c *Cache) Delte(key string) {
	delete(c.data, key)
}
