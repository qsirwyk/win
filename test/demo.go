package main

import (
	"fmt"
	"github.com/qsirwyk/win"
	"time"
)

func main() {
	for {
		fmt.Println(win.GetAsyncKeyState(1))
		time.Sleep(time.Millisecond * 300)
	}
}
