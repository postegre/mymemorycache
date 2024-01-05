package main

import (
	"fmt"
	"time"

	"github.com/your-username/cache"
)

func main() {
	// Создаем новый экземпляр кеша
	myCache := cache.NewCache()

	// Устанавливаем значение в кеш
	myCache.Set("key1", "value1")
	myCache.Set("key2", "value2")

	// Получаем значение из кеша
	value := myCache.Get("key1")
	fmt.Println("Value for key 'key1':", value)

	// Удаляем значение из кеша
	myCache.Delete("key2")

	// Пытаемся получить удаленное значение
	value = myCache.Get("key2")
	fmt.Println("Value for key 'key2':", value)

	// Добавляем значение с таймаутом
	myCache.Set("key3", "value3 with timeout", 2*time.Second)

	// Ждем таймаут
	time.Sleep(3 * time.Second)

	// Пытаемся получить просроченное значение
	value = myCache.Get("key3")
	fmt.Println("Value for key 'key3' after timeout:", value)
}
