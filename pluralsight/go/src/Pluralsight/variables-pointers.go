package main

import (
	"fmt"
	"os"
)

func main() {
	name := os.Getenv("USER")
	course := "Docker Deep Dive"

	fmt.Println("\nHi", name, "you're currently watching", course)

	changeCourse(&course)

	fmt.Println("\nYou are now watching course", course)

}

func changeCourse(course *string) string {
	*course = "First Look: Native Docker Clistering"
	fmt.Println("\nTrying to change your course to", *course)
	return *course
}
