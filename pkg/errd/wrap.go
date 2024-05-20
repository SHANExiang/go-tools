package main

import (
	"errors"
	"fmt"
	"log"
)

func Wrap(err *error, f string, v ...interface{}) {
	if *err != nil {
		*err = fmt.Errorf(f+": %w", append(v, *err)...)
	}
}

func testWrap() (err error) {
	defer Wrap(&err, "failed to process")

	err = errors.New("test error")
	return fmt.Errorf("failed err: %w", err)
}

func main() {
    log.Println(testWrap())
}
