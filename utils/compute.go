package utils

import (
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func ComputeIPsArray(ips string, numIPs int) ([]string, error) {
	ips = strings.TrimSpace(ips)

	ipsArray := strings.Split(ips, ",")
	for i, ip := range ipsArray {
		ipsArray[i] = strings.TrimSpace(ip)
	}

	if numIPs == 1 || len(ipsArray) > 1 {
		return ipsArray, nil
	}

	if numIPs > 1 {
		baseIP := net.ParseIP(ipsArray[0])
		if baseIP == nil {
			return nil, fmt.Errorf("invalid IP address: %s", ipsArray[0])
		}

		for i := 1; i < numIPs; i++ {
			nextIP := incrementIP(baseIP, i)
			if nextIP == nil {
				return nil, fmt.Errorf("IP overflow when generating IP #%d", i+1)
			}
			ipsArray = append(ipsArray, nextIP.String())
		}
		return ipsArray, nil
	}

	return nil, fmt.Errorf("generic error computing IPs")
}

func incrementIP(ip net.IP, increment int) net.IP {
	ip4 := ip.To4()
	if ip4 == nil {
		return nil
	}

	ipInt := uint32(ip4[0])<<24 + uint32(ip4[1])<<16 + uint32(ip4[2])<<8 + uint32(ip4[3])
	ipInt += uint32(increment)

	if ipInt < uint32(ip4[0])<<24+uint32(ip4[1])<<16+uint32(ip4[2])<<8+uint32(ip4[3]) {
		return nil
	}

	newIP := net.IPv4(byte(ipInt>>24), byte(ipInt>>16), byte(ipInt>>8), byte(ipInt))
	return newIP
}

func ComputeMountPointsArray(mountPoints string, numDisks int) ([]string, error) {
	mountPoints = strings.TrimSpace(mountPoints)
	if mountPoints == "" {
		mountPoints = "/data/agent00"
	}

	mountPointsArray := strings.Split(mountPoints, ",")
	for i, mp := range mountPointsArray {
		mountPointsArray[i] = strings.TrimSpace(mp)
	}

	if len(mountPointsArray) > 1 && len(mountPointsArray) == numDisks {
		return mountPointsArray, nil
	}

	baseMountPoint := mountPointsArray[0]
	if len(baseMountPoint) == 0 {
		return nil, fmt.Errorf("empty mount point")
	}

	suffixStart := len(baseMountPoint)
	for i := len(baseMountPoint) - 1; i >= 0; i-- {
		if !unicode.IsDigit(rune(baseMountPoint[i])) {
			break
		}
		suffixStart = i
	}

	base := baseMountPoint[:suffixStart]
	suffix := baseMountPoint[suffixStart:]

	if suffix == "" {
		return nil, fmt.Errorf("cannot determine numeric suffix in mount point")
	}

	start, err := strconv.Atoi(suffix)
	if err != nil {
		return nil, fmt.Errorf("invalid numeric suffix: %v", err)
	}

	width := len(suffix)
	mountPointsArray = []string{}
	for i := 0; i < numDisks; i++ {
		mp := fmt.Sprintf("%s%0*d", base, width, start+i)
		mountPointsArray = append(mountPointsArray, mp)
	}

	return mountPointsArray, nil
}

func ComputeDisksArray(disks string, numDisks int) ([]string, error) {
	disks = strings.TrimSpace(disks)
	if disks == "" {
		disks = "/dev/sda"
	}

	if numDisks < 1 {
		return nil, fmt.Errorf("numDisks must be >= 1")
	}

	disksArray := strings.Split(disks, ",")
	for i, d := range disksArray {
		disksArray[i] = strings.TrimSpace(d)
	}
	if len(disksArray) > 1 && len(disksArray) == numDisks {
		return disksArray, nil
	}

	base := disks[:len(disks)-1]
	suffix := disks[len(disks)-1]

	if suffix < 'a' || suffix > 'z' {
		return nil, fmt.Errorf("unsupported disk suffix: %c", suffix)
	}

	disksArray = []string{}
	for i := 0; i < numDisks; i++ {
		letter := suffix + byte(i)
		if letter > 'z' {
			return nil, fmt.Errorf("disk suffix overflow")
		}
		disksArray = append(disksArray, fmt.Sprintf("%s%c", base, letter))
	}

	return disksArray, nil
}

func ComputePortsArray(ports string, numPorts int) ([]int, error) {
	ports = strings.TrimSpace(ports)
	portsArray := strings.Split(ports, ",")

	var result []int
	for _, port := range portsArray {
		port = strings.TrimSpace(port)
		if port == "" {
			continue
		}
		p, err := strconv.Atoi(port)
		if err != nil {
			return nil, fmt.Errorf("invalid port: %s", port)
		}
		result = append(result, p)
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("no valid ports provided")
	}

	if len(result) == 1 && numPorts > 1 {
		basePort := result[0]
		for i := 1; i < numPorts; i++ {
			result = append(result, basePort+i)
		}
	}

	if len(result) != numPorts {
		return nil, fmt.Errorf("port count mismatch: expected %d, got %d", numPorts, len(result))
	}

	return result, nil
}

func ComputeNamesArray(names string, numNodes int) []string {
	rand.Seed(time.Now().UnixNano())

	names = strings.TrimSpace(names)

	if names == "" {
		names = fmt.Sprintf("node-%d", rand.Intn(1000))
	}

	if numNodes == 1 {
		return []string{names}
	}

	namesArray := strings.Split(names, ",")
	for i, name := range namesArray {
		namesArray[i] = strings.TrimSpace(name)
	}

	if len(namesArray) == numNodes {
		return namesArray
	}

	baseName := namesArray[0]
	result := make([]string, numNodes)
	for i := 0; i < numNodes; i++ {
		result[i] = fmt.Sprintf("%s-%d", baseName, rand.Intn(1000))
	}

	return result
}
