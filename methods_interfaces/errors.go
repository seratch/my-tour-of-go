package main

import (
	"fmt"
	"strconv"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}

	/*
		// Atoi returns the result of ParseInt(s, 10, 0) converted to type int.
		func Atoi(s string) (int, error) {
			const fnAtoi = "Atoi"
			i64, err := ParseInt(s, 10, 0)
			if nerr, ok := err.(*NumError); ok {
				nerr.Func = fnAtoi
			}
			return int(i64), err
		}
	*/
	i, err := strconv.Atoi("42a")
	if err != nil {
		// couldn't convert number: strconv.Atoi: parsing "42a": invalid syntax
		fmt.Printf("couldn't convert number: %v\n", err)
		return
	}
	fmt.Println("Converted integer:", i)

}
