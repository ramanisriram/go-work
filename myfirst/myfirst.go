package main

import (
	"fmt"

	"github.com/sriram1113/go-work/myfirst/morestrings"
	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println(morestrings.AddStr("Sri"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))

}

