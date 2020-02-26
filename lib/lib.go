package lib

import (
	"fmt"
	"os"
	"sync"
)

var lock sync.Mutex

func Now(name string) {
	f, err := os.OpenFile("../out/test.txt", os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	for i:=0; i<100; i++ {
		fmt.Fprintln(f, name)
	}
	f.Close()
}