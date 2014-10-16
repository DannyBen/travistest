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

	exists, err := travistest.Exists(handler.Filename(key))
	if err != nil {
		fmt.Println("10 ERR >", err.Error())
	} else {
		fmt.Println("10 OK  >")
	}

	// r := handler.Get(key)
	// fmt.Println("10 --->", string(r))
	// fmt.Println(travistest.Exists(handler.Filename(key)))

	if exists {
		fi, err := os.Stat(handler.Filename(key))
		if err != nil {
			fmt.Println("20 ERR >", err.Error())
		} else {
			fmt.Println("20 OK  >", fi.Mode())
		}
	} else {
		fmt.Println("25 ERR > DOES NOT EXIST")
	}

	// file := "./travis.go"
	// fi, err = os.Stat(file)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("2 --->", fi.Mode())

	// Output:
	// ------ > RESULT
}
