package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"golang.org/x/sync/singleflight"
)

// 1. A single 'Group' is created to manage coalescing.
// This is typically a global or part of a service struct.
var requestGroup singleflight.Group

// 2. This is our slow, expensive function (e.g., a DB or API call).
func fetchFromDatabase(key string) (interface{}, error) {
	log.Printf("Fetching data for key [%s] from database... (THIS SHOULD ONLY APPEAR ONCE!)", key)
	// Simulate a slow network call
	time.Sleep(2 * time.Second)
	
	// The data is fetched
	data := fmt.Sprintf("Data for %s", key)
	
	return data, nil
}

// 3. This is our handler function that gets all the concurrent requests.
func handleRequest(key string) (interface{}, error) {
	// The magic is here: requestGroup.Do(key, fn)
	// - 'key' is the unique identifier for the resource.
	// - 'fn' is the function to call *only once*.
	//
	// How it works:
	// - The FIRST goroutine to call Do() with a specific key will execute fn().
	// - ANY OTHER goroutine that calls Do() with the *same key* while fn()
	//   is still running will block and wait.
	// - When fn() completes, its result (data, err) is returned to the
	//   first goroutine *and* all the waiting goroutines.
	//
	// - The 'shared' boolean indicates if the result was shared (i.e., you
	//   were a "waiter" and not the "fetcher").
	
	v, err, shared := requestGroup.Do(key, func() (interface{}, error) {
		// This anonymous function is the work we want to coalesce.
		return fetchFromDatabase(key)
	})

	log.Printf("Request for [%s] completed. Shared: %v", key, shared)
	return v, err
}

func main() {
	var wg sync.WaitGroup
	
	// The "hot key" we're all going to request at the same time.
	hotKey := "product-123"

	// Simulate 50,000 requests (we'll use 50 for this demo)
	numRequests := 50
	wg.Add(numRequests)

	log.Println("Simulating 50 concurrent requests for the SAME key...")

	for i := 0; i < numRequests; i++ {
		go func(requestID int) {
			defer wg.Done()
			
			log.Printf("Request %d: Firing request for key %s", requestID, hotKey)
			
			data, err := handleRequest(hotKey)
			
			if err != nil {
				log.Printf("Request %d: Failed: %v", requestID, err)
			} else {
				log.Printf("Request %d: Got data: %s", requestID, data)
			}
		}(i)
	}

	wg.Wait()
	log.Println("All requests complete.")
}