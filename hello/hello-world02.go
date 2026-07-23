package hello

import (
	"fmt"
	"os"
)

func Hello_World02_Example() {
	name := "Go"
	fmt.Println("Hello,", name)
	fmt.Println("Args:", os.Args)
}
