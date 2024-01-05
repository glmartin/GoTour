package main

import (
	"fmt"
)

type IPAddr [4]byte

func (ip IPAddr) String() string {
	ipStr := ""
	for _, x := range ip {
		if len(ipStr) > 0 {
			ipStr += "."
		}
		ipStr += fmt.Sprint(x)
	}
	return ipStr
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
