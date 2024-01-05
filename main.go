package main

import "fmt"

func Hellp(name string) string {
	return "Hellp, " + name + "!"
}

func main() {
	fmt.Println(Hellp("Josh"))
}
