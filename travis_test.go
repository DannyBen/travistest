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
	fmt.Println("10 --->", string(r))
	// fmt.Println(travistest.Exists(handler.Filename(key)))

	yes, err := travistest.Exists(handler.Filename(key))
	if err != nil {
		fmt.Println("60 ERR>", err.Error())
	}
	if yes {
		fmt.Println("20 ---> EXISTS")
		fi, err := os.Stat(handler.Filename(key))
		if err != nil {
			fmt.Println("50 ERR>", err.Error())
		}
		fmt.Println("30 --->", fi.Mode())
	} else {
		fmt.Println("40 ---> DOES NOT EXIST")
	}

	// file := "./travis.go"
	// fi, err = os.Stat(file)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("2 --->", fi.Mode())

	// Output:
	// ------> RESULT
}
