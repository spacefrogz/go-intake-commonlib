package main

import (
	"fmt"

	"github.com/spacefrogz/go-intake-commonlib/errors"
)

func main() {
	e := errors.SE("0001")
	fmt.Println(e)
}
