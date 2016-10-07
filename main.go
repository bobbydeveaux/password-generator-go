package main

import (
	"flag"
	"fmt"
	"password-generator/generator"
)

var (
	length       = flag.Int("length", 20, "Length of password")
	useUpperCase = flag.Bool("upper", false, "Use upper-case characters A-Z")
	useLowerCase = flag.Bool("lower", false, "Use lower-case characters a-z")
	useNumbers   = flag.Bool("numbers", false, "Use numbers 0-9")
	useSpecial   = flag.Bool("special", false, "Use special characters $%&*@^")
)

func main() {
	flag.Parse()

	password, error := generator.NewPassword(
		*length,
		*useUpperCase,
		*useLowerCase,
		*useNumbers,
		*useSpecial,
	)

	if error != nil {
		fmt.Println("Error:", error)
		return
	}

	fmt.Println(password)
}
