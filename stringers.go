package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string {
	var final_string string

	for _, v := range ip {
		final_string = fmt.Sprintf("%v%v.", final_string, v)
	}

	return final_string
}

func main() {
	hosts := IPAddr{1, 2, 4, 5}

	fmt.Print(hosts)
}
