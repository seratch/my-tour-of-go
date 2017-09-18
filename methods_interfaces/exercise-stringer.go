package main

import (
	"fmt"
	//"strconv"
	//"strings"
)

type IPAddr [4]byte

// Add a "String() string" method to IPAddr.
func (ipAddr IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ipAddr[0], ipAddr[1], ipAddr[2], ipAddr[3])

	//strArr := make([]string, 4)
	//for idx, element := range iaddr {
	//	strArr[idx] = strconv.Itoa(int(element))
	//}
	//return strings.Join(strArr, ".")
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
