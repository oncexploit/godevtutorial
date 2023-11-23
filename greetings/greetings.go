package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	//If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// Create a message using a random format.
	message := fmt.Sprintf(randomFormat(), name)
	//missing argument name so that the test can not pass.
	//message := fmt.Sprint(randomFormat())
	return message, nil
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	//A map to associate name with strings.
	messages := make(map[string]string)
	//Loop through the received slice of names,
	//calling the hello function to get a message for each name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		//In the map, associate the retrieved message with the name.
		messages[name] = message
	}

	return messages, nil
}

//Hellos()相比于Hello()，从处理单一输入变为处理多值输入。然而直接对Hello函数进行修改，
//那么已经使用Hello编写代码的程序将被破坏。
//Hellos的作用是既能处理多参数，又保留了旧功能，实现向后的兼容性。

func randomFormat() string {
	//A slice of message formats
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
