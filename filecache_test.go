package filecache_test

import (
	"fmt"
	filecache "github.com/DannyBen/travistest"
)

func Example() {
	// Get a handler and set a directory + 1 hour cache life
	handler := filecache.Handler{}

	// Data to store in cache
	data := []byte("Joey doesn't share food")

	// Store the data in the cache
	handler.Set("testkey", data)

	// Retrieve the data from the cache
	r := handler.Get("testkey")

	// Show the result
	fmt.Println(string(r))

	// Wait for some seconds
	// time.Sleep(7 * time.Second)

	// By now the cache is invalid
	// r = handler.Get("testkey")
	// if r == nil {
	// 	fmt.Println("Cache expired")
	// }

	// Output:
	// Joey doesn't share food
}
