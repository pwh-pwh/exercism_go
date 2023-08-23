package airportrobot

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
type Greeter interface {
	LanguageName() string
	Greet(string) string
}

func SayHello(name string, greeter Greeter) string {
	// => "I can speak German: Hallo Dietrich!"
	return fmt.Sprintf("I can speak %s: %s", greeter.LanguageName(), greeter.Greet(name))
}

// Italian
type Italian struct {
}

func (i Italian) LanguageName() string {
	return "Italian"
}

func (i Italian) Greet(s string) string {
	return fmt.Sprintf("Ciao %s!", s)
}

type Portuguese struct {
}

func (p Portuguese) LanguageName() string {
	return "Portuguese"
}

func (p Portuguese) Greet(s string) string {
	return fmt.Sprintf("Olá %s!", s)
}
