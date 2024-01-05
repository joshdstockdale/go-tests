package main

import "fmt"

const englishHellpPrefix = "Hellp, "

func Hellp(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHellpPrefix + name + "!"
}

func main() {
	fmt.Println(Hellp("Josh"))
}
