package travistest_test

import (
	"fmt"
	"github.com/DannyBen/travistest"
)

func Example() {
	handler := travistest.Handler{"./cache", 1}
	data := []byte("HELLO")
	key := "key"
	handler.Set(key, data)
	r := handler.Get(key)
	fmt.Println(string(r))
	fmt.Println(travistest.Exists(handler.Filename("testkey")))

	// Output:
	// HELLO
	// true <nil>
}
