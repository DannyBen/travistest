package travistest_test

import (
	"fmt"
	"github.com/DannyBen/travistest"
	"os"
)

func Example() {
	handler := travistest.Handler{"./cache", 1}
	data := []byte("HELLO")
	key := "key"
	handler.Set(key, data)
	r := handler.Get(key)
	fmt.Println("0 --->", string(r))
	// fmt.Println(travistest.Exists(handler.Filename(key)))

	fi, err := os.Stat(handler.Filename(key))
	fmt.Println("1 --->", fi.Mode())

	file := "./travis.go"
	fi, err = os.Stat(file)
	if err != nil {
		panic(err)
	}
	fmt.Println("2 --->", fi.Mode())

	// Output:
	// HELLO
	// true <nil>
}
