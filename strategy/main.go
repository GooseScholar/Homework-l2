package main

import (
	"fmt"
	"sync"
)

//Интерфейс стратегии
type evictionAlgo interface {
	evict(c *cache)
}

//Первым пришел, первым ушел
type fifo struct {
}

func (l *fifo) evict(c *cache) {
	fmt.Println("Evicting by fifo strtegy")
}

//Наиболее давно использовавшиеся
type lru struct {
}

func (l *lru) evict(c *cache) {
	fmt.Println("Evicting by lru strtegy")
}

//Наименее часто использовавшиеся
type lfu struct {
}

func (l *lfu) evict(c *cache) {
	fmt.Println("Evicting by lfu strtegy")
}

//Кеш
type cache struct {
	mx           sync.RWMutex
	storage      map[string]string
	evictionAlgo evictionAlgo
	capacity     int
	maxCapacity  int
}

func initCache(e evictionAlgo) *cache {
	storage := make(map[string]string)
	return &cache{
		storage:      storage,
		evictionAlgo: e,
		capacity:     0,
		maxCapacity:  2,
	}
}

func (c *cache) setEvictionAlgo(e evictionAlgo) {
	c.evictionAlgo = e
}

func (c *cache) add(key, value string) {
	c.mx.Lock()
	defer c.mx.Unlock()
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *cache) get(key string) {
	c.mx.RLock()
	defer c.mx.RUnlock()
	delete(c.storage, key)
}

func (c *cache) evict() {
	c.evictionAlgo.evict(c)
	c.capacity--
}

func main() {
	lfu := &lfu{}
	cache := initCache(lfu)

	cache.add("a", "1")
	cache.add("b", "2")
	cache.add("c", "3")
	fmt.Println(cache)

	lru := &lru{}
	cache.setEvictionAlgo(lru)
	cache.add("d", "4")
	fmt.Println(cache)

	fifo := &fifo{}
	cache.setEvictionAlgo(fifo)
	cache.add("e", "5")
	fmt.Println(cache)
}
