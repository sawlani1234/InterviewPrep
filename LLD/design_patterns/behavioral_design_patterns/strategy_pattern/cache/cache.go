package cache

import evictionstrategy "strategy_pattern/eviction_strategy"


type Cache struct {
   
   cacheEvictionStrategy evictionstrategy.EvictionStrategy
   mp map[int]int
}

func NewCache(e evictionstrategy.EvictionStrategy) *Cache {

	return &Cache{cacheEvictionStrategy: e,mp : make(map[int]int,0)}
	
}

func (c *Cache) Add(key int,val int) {
   c.mp[key] = val
}

func (c *Cache) Remove(key int) {
	delete(c.mp,key)
}

func (c *Cache) Evict() {
   c.cacheEvictionStrategy.Evict()
}