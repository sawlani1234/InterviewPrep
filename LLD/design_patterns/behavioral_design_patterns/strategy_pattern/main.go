package main

import evictionstrategy "strategy_pattern/eviction_strategy"
import cache "strategy_pattern/cache"

/*
Notes About Strategy pattern
1. Strategy pattern basically defines a family of algorithms
2. all algorithms implement an interface 
3. the clients injects the strategy it needs to use and call the appopriate strategy 

*/
func main() {
	sqlDb := cache.NewCache(evictionstrategy.Lmu{})
	sqlDb.Evict()
}