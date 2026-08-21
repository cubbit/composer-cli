package utils

import (
	"fmt"
	"net"
	"regexp"
	"strings"
)

var nodeNameRegex = regexp.MustCompile(`^[a-z0-9][a-z0-9\-\.]+[a-z0-9]$`)

func IsValidIP(ip string) bool {
	return net.ParseIP(ip) != nil
}

func ValidateIPsInput(ips string, numNodes int) error {
	ipList := strings.Split(ips, ",")
	for _, ip := range ipList {
		if net.ParseIP(strings.TrimSpace(ip)) == nil {
			return fmt.Errorf("invalid IP address: %s", ip)
		}
	}

	if len(ipList) > 1 && len(ipList) != numNodes {
		return fmt.Errorf("IP count mismatch")
	}

	return nil
}

func ValidateNamesInput(names string, numNodes int) error {
	if names == "" {
		return nil
	}

	nameList := strings.Split(names, ",")
	nameSet := make(map[string]bool)

	for _, name := range nameList {
		name = strings.TrimSpace(name)
		if !nodeNameRegex.MatchString(name) {
			return fmt.Errorf("invalid name format: %s", name)
		}
		if nameSet[name] {
			return fmt.Errorf("duplicate name: %s", name)
		}
		nameSet[name] = true
	}

	if len(nameList) > 1 && len(nameList) != numNodes {
		return fmt.Errorf("name count mismatch")
	}

	return nil
}
