package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	// => Welcome to my party, Christiane!
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	// => Happy birthday Frank! You are now 58 years old!
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	// =>
	// Welcome to my party, Christiane!
	// You have been assigned to table 027. Your table is on the left, exactly 23.8 meters from here.
	// You will be sitting next to Frank.
	return fmt.Sprintf("Welcome to my party, %s!\n"+
		"You have been assigned to table %03d. Your table is %s, exactly %.1f meters from here.\n"+
		"You will be sitting next to %s.", name, table, direction, distance, neighbor)
}
