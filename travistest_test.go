package travistest_test

import (
	"fmt"
	"github.com/DannyBen/filecache"
	"github.com/DannyBen/travistest"
)

func Example() {
	// travistest.CreateFile("file")
	// yes, err := travistest.Exists("file")
	// fmt.Println(yes, err)

	handler := filecache.Handler{"./cache", 1}
	data := []byte("HELLO")
	handler.Set("testkey", data)
	r := handler.Get("testkey")
	fmt.Println(string(r))
	fmt.Println(travistest.Exists(handler.Filename("testkey")))

	// Output:
	// HELLO
	// true <nil>
}
