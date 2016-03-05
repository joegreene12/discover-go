package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, we are Holberton School")
	p := fmt.Println
	now := time.Now()
	p("the date is", now)
	p("the year is", now.Year())
	p("the month is", now.Month())
	p("the day is", now.Day())

}
