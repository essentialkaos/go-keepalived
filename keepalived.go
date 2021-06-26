// Package keepalived is a simple Go package for reading virtual IP info from keepalived config
package keepalived

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2021 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bufio"
	"errors"
	"net"
	"os"
	"regexp"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// ipValidationRegex simple IPv4 validation regexp
var ipValidationRegex = regexp.MustCompile(`^[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}$`)

// ////////////////////////////////////////////////////////////////////////////////// //

// GetVirtualIP returns virtual IP from keepalived configuration file
func GetVirtualIP(config string) (string, error) {
	fd, err := os.OpenFile(config, os.O_RDONLY, 0)

	if err != nil {
		return "", err
	}

	defer fd.Close()

	r := bufio.NewReader(fd)
	s := bufio.NewScanner(r)

	var vrrpFound, virtualFound bool

	for s.Scan() {
		line := s.Text()

		if strings.Contains(line, "vrrp_instance") {
			vrrpFound = true
		}

		if !vrrpFound {
			continue
		}

		if strings.Contains(line, "virtual_ipaddress") {
			virtualFound = true
		}

		if !virtualFound {
			continue
		}

		if !strings.Contains(line, " label ") {
			continue
		}

		ip := extractIP(line)

		if ip == "" {
			return "", errors.New("Can't find virtual IP info")
		}

		return ip, nil
	}

	return "", errors.New("Can't find virtual IP info")
}

// IsMaster returns true if given virtual IP is used on this server
func IsMaster(virtualIP string) (bool, error) {
	if !ipValidationRegex.MatchString(virtualIP) {
		return false, errors.New("Given IP is not valid IPv4 address")
	}

	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return false, err
	}

	for _, addr := range addrs {
		if addr.String() == virtualIP+"/32" {
			return true, nil
		}
	}

	return false, nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// extractIP extracts IP from the address defenition
func extractIP(data string) string {
	data = strings.TrimSpace(data)
	index := strings.Index(data, " ")

	if index == -1 {
		return ""
	}

	ip := data[:index]

	if !ipValidationRegex.MatchString(ip) {
		return ""
	}

	return ip
}
