package keepalived

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2019 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
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
