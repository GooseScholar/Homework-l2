package cache

import "sync"

//Cache хранилище дат
type Cache struct {
	mx   sync.RWMutex
	Data map[string]map[string]struct{}
}

//NewCache cоздание кеша
func NewCache() *Cache {
	return &Cache{
		Data: make(map[string]map[string]struct{}),
	}
}

//PutEvent запись данных в кеш
func (c *Cache) PutEvent(d string, id string) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.Data[d] = make(map[string]struct{})
	c.Data[d][id] = struct{}{}
}

//GetEvents получение данных из кеша
func (c *Cache) GetEvents(d string, id string) (event string, b bool) {
	c.mx.RLock()
	defer c.mx.RUnlock()
	_, b = c.Data[d][id]
	return
}

//DeleteEvent удаление данных из кеша
func (c *Cache) DeleteEvent(d string, id string) {
	c.mx.Lock()
	defer c.mx.Unlock()
	delete(c.Data[d], id)
}

//UpdateEvent обновление событый
func (c *Cache) UpdateEvent(d string, nd string, id string) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.Data[nd] = make(map[string]struct{})
	c.Data[nd][id] = struct{}{}
	delete(c.Data[d], id)
}
