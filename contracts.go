package contracts

import (
	"fmt"
	"log"
)

/*
 * Contracts on functions consist of preconditions and postconditions.
 */

type fn func(string)

func require(f fn, s string, desc string, expression bool) bool {

	fmt.Println("require")
	fmt.Println(desc)

	if !expression {
		log.Fatal("Aaaa!")
	}
	/* example of expression */
	if s == "" {
		log.Fatal("Aaaaa!")
	}
	return true
}

func ensure(f fn, s string, desc string, expression bool) func(s string) {
	fmt.Println("ensure")
	fmt.Println(desc)

	return func(s string) {
		f(s)
	}
}

func Contract(f fn, s string, desc string, expression bool) func(s string) {
	fmt.Println("contract")
	require(f, s, desc, expression)

	return ensure(f, s, desc, expression)
}
