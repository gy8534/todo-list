package main

import (
	"flag"
	"fmt"
)

type Task struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Body        string `json:"body"`
	IsCompleted bool   `json:"is_completed"`
}

var tasks []Task

func main() {
	list := flag.Bool("l", false, "list all the tasks")
	add := flag.Bool("a", false, "add a task to the list")
	flag.Parse()

	fmt.Print(*list)

}
