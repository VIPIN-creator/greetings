package greetings

import (
    "errors"
    "fmt"
    "math/rand"
)

func Hellos(names []string) (map[string]string, error) {
    messages := make(map[string]string)

    //In this for loop, range returns two values: the index of the current item in the loop and a copy of the item's value. 
    // You don't need the index, so you use the Go blank identifier (an underscore) to ignore it. 
    for _, name := range names {
       
        message, err := Hello(name)

        if err != nil {
            return nil, err
        }

        messages[name] = message
    }

    return messages, nil
}

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {

    if name == "" {
        return "", errors.New("empty name")
    }

    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf(randFormat(), name)
    return message, nil
}

func randFormat() string {

    // A slice of strings. No size mentioned in [], so Go compiler knows the size of slice is dynamic
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }

    return formats[rand.Intn(len(formats))]
}