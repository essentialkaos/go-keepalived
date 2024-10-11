// Package keepalived is a simple Go package for reading virtual IP info from keepalived config
package keepalived

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	MASTER = "master" // MASTER is keepalived master state
	BACKUP = "backup" // BACKUP is keepalived backup state
)

// ////////////////////////////////////////////////////////////////////////////////// //

// State is keepalived state
type State string

// Instance contains basic information about keepalived instance
type Instance struct {
	Name      string
	Interface string
	State     State
	Priority  int
	Addr      *Address
}

// Address contains info about virtual IP
type Address struct {
	IP  net.IP
	Net *net.IPNet
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Config is path to keepalived configuration file
var Config = "/etc/keepalived/keepalived.conf"

// ////////////////////////////////////////////////////////////////////////////////// //

// GetByName returns keepalived instance with given name
func GetByName(name string) (*Instance, error) {
	instances, err := List()

	if err != nil {
		return nil, err
	}

	for _, instance := range instances {
		if instance.Name == name {
			return instance, nil
		}
	}

	return nil, nil
}

// GetByInterface returns keepalived instance for given interface name
func GetByInterface(interfaceName string) (*Instance, error) {
	instances, err := List()

	if err != nil {
		return nil, err
	}

	for _, instance := range instances {
		if instance.Interface == interfaceName {
			return instance, nil
		}
	}

	return nil, nil
}

// List returns slice with keepalived instances
func List() ([]*Instance, error) {
	fd, err := os.OpenFile(Config, os.O_RDONLY, 0)

	if err != nil {
		return nil, err
	}

	defer fd.Close()

	s := bufio.NewScanner(fd)

	var result []*Instance
	var instance *Instance
	var addressFound bool

	for s.Scan() {
		line := strings.Trim(s.Text(), " \t")

		switch {
		case strings.HasPrefix(line, "vrrp_instance "):
			if instance != nil {
				result = append(result, instance)
			}

			instance = &Instance{Name: readField(line, 1)}
			continue

		case strings.HasPrefix(line, "virtual_ipaddress "):
			addressFound = true
			continue

		case strings.HasPrefix(line, "state "):
			instance.State = State(strings.ToLower(readField(line, 1)))
			continue

		case strings.HasPrefix(line, "interface "):
			instance.Interface = readField(line, 1)
			continue

		case strings.HasPrefix(line, "priority "):
			priority, err := strconv.Atoi(readField(line, 1))

			if err != nil {
				return nil, fmt.Errorf("Can't parse priority: %w", err)
			}

			instance.Priority = priority
			continue
		}

		if !addressFound {
			continue
		}

		instance.Addr, err = extractAddr(readField(line, 0))

		if err != nil {
			return nil, fmt.Errorf("Can't parse virtual address info: %w", err)
		} else {
			addressFound = false
		}
	}

	if instance != nil {
		result = append(result, instance)
	}

	return result, nil
}

// IsInstaceMaster returns true if keepalived instance with given name is in MASTER state
func IsInstaceMaster(instanceName string) (bool, error) {
	instance, err := GetByName(instanceName)

	if err != nil {
		return false, err
	}

	if instance == nil {
		return false, fmt.Errorf("There is no instance with name %q", instanceName)
	}
	interfaces, err := net.Interfaces()

	if err != nil {
		return false, fmt.Errorf("Can't get list of interfaces: %w", err)
	}

	for _, intrf := range interfaces {
		addrs, err := intrf.Addrs()

		if err != nil {
			return false, fmt.Errorf("Can't get list of interface %q address: %w", intrf.Name, err)
		}

		for _, addr := range addrs {
			if addr.String() == instance.Addr.IP.String()+"/32" {
				return true, nil
			}
		}
	}

	return false, nil
}

// IsMaster returns true if any of keepalived instances is in MASTER state
func IsMaster() (bool, error) {
	instances, err := List()

	if err != nil {
		return false, err
	}

	if len(instances) == 0 {
		return false, nil
	}

	interfaces, err := net.Interfaces()

	if err != nil {
		return false, fmt.Errorf("Can't get list of interfaces: %w", err)
	}

	for _, intrf := range interfaces {
		addrs, err := intrf.Addrs()

		if err != nil {
			return false, fmt.Errorf("Can't get list of interface %q address: %w", intrf.Name, err)
		}

		for _, addr := range addrs {
			for _, instance := range instances {
				if addr.String() == instance.Addr.IP.String()+"/32" {
					return true, nil
				}
			}
		}
	}

	return false, nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// String returns string representation of virtual IP
func (addr *Address) String() string {
	if addr == nil || addr.IP == nil {
		return ""
	}

	if addr.Net == nil {
		return addr.IP.String()
	}

	netSize, _ := addr.Net.Mask.Size()

	return fmt.Sprintf("%s/%d", addr.IP.String(), netSize)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// extractAddr extracts virtual address info
func extractAddr(data string) (*Address, error) {
	if strings.Contains(data, "/") {
		vIP, vNet, err := net.ParseCIDR(data)

		if err != nil {
			return nil, err
		}

		return &Address{IP: vIP, Net: vNet}, nil
	}

	vIP := net.ParseIP(data)

	if vIP == nil {
		return nil, fmt.Errorf("Invalid IP")
	}

	return &Address{IP: vIP}, nil
}

// readField reads field with given index from data
func readField(data string, index int) string {
	if data == "" || index < 0 {
		return ""
	}

	curIndex, startPointer := -1, -1

MAINLOOP:
	for i, r := range data {
		if r == ' ' {
			if curIndex == index {
				return data[startPointer:i]
			}

			startPointer = -1
			continue MAINLOOP
		}

		if startPointer == -1 {
			startPointer = i
			curIndex++
		}
	}

	if index > curIndex {
		return ""
	}

	return data[startPointer:]
}
