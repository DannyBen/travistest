package travistest_test

import (
	"fmt"
	"github.com/DannyBen/travistest"
)

func Example() {
	travistest.CreateFile("file")
	yes, err := travistest.Exists("file")
	fmt.Println(yes, err)
	// Output:
	// true <nil>
}
