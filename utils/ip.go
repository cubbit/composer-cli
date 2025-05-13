package utils

import "net"

func IsValidIP(ip string) bool {
	return net.ParseIP(ip) != nil
}
