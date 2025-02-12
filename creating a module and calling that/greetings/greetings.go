package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	// if no name was given, return an error with a message
	if name == "" {
		return name, errors.New("empyt name")
	}
	// creating a mesage using a random format
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil

}

// randomFormat returns one of a set of greeting messages. The returned message is selected at random.
// message is seleced at random

func randomFormat() string {
	// A slice of message formats.

	formats := []string{
		"hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	// Return a randomly selected message format by specifying a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}

// Hello retuns a map that asscoicates each of the named people with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	// A map o associate names with messages.
	messages := make(map[string]string)

	// Loop through the received slice of names, calling the Hello function to get a message for each name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		// In the map, associate the retrieved message with the name.
		messages[name] = message
	}
	return messages, nil
}
