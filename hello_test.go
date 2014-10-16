package main_test

import (
	"fmt"
	"io/ioutil"
	"os"
	// "syscall"
	"time"
)

func Example() {
	file := os.TempDir() + "/somefile.txt"
	fmt.Println("Started   :", time.Now())
	ioutil.WriteFile(file, []byte("hello"), 0666)
	// syscall.Sync()
	fi, _ := os.Stat(file)
	fmt.Println("ModTime   :", fi.ModTime())
	fmt.Println("Now       :", time.Now())
	fmt.Println("Time Since:", time.Since(fi.ModTime()).Seconds())

	// OUTPUT:
	// Expecting something different
}
