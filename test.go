package main

import "fmt"

// TODO:  is a comment

/* Another comment */

type Bar struct {
	i int
	b bool
}

func Foo() *Bar {
	return new(Bar)
}

func Baz() error {
	return nil
}

func main() {
	if b := Baz(); b != nil {
		fmt.Println("uh oh!")
	}

	n := []string{}
	n = append(n, "foo")

	m := map[string]int{
		"foo": 1,
	}
	if m == nil {
		fmt.Println("uh oh!")
	}

	rs := `sdlfkjsldkfjlsdkjf
	lksdjflksdjflksdjf`

	fmt.Println(rs)

	c := make(chan string)
	c <- "foo"

}
