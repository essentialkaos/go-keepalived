package keepalived

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleGetVirtualIP() {
	ip, err := GetVirtualIP("/etc/keepalived/keepalived.conf")

	fmt.Printf("Virtual IP: %s\n", ip)
	fmt.Printf("Error: %v\n", err)
}

func ExampleIsMaster() {
	isMaster, err := IsMaster("191.12.11.18")

	fmt.Printf("Is Master: %t\n", isMaster)
	fmt.Printf("Error: %v\n", err)
}
