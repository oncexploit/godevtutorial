package main

import (
	"fmt"
	"github.com/greetings"
	"log"
)

func main() {
	//设置日志条目前缀和flag
	log.SetPrefix("greetings:")
	log.SetFlags(0)

	//A slice of names.
	names := []string{"Gladys", "Robert", "Darwin"}

	//request a greeting message
	message, err := greetings.Hellos(names)
	// if an error was returned, print it to the console and exit.
	if err != nil {
		log.Fatal(err)
	}

	//if no error was returned, print the returned message.
	fmt.Println(message)
}
