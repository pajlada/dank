package main

import "fmt"

func main() {
	fmt.Println(fmt.Sprintf("do you like that fixmie %s"))
}

func getCoolError() error {
	return errors.New(fmt.Sprintf("something was wrong.... %s", lol))
}
