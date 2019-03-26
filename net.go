package gotools

import (
	"net"
	"strconv"
	"strings"
)

type IPVersion int

const (
	VersionUnknown IPVersion = iota
	IPv4           IPVersion = iota
	IPv6           IPVersion = iota
)

func ParseIPVersion(str string) IPVersion {
	ip := net.ParseIP(str)
	if ip == nil {
		return VersionUnknown
	}
	if strings.Contains(str, ":") {
		return IPv6
	}
	return IPv4
}

func IPv4ByIfaceName(iface string) (ips []net.IP, err error) {
	var (
		i     *net.Interface
		addrs []net.Addr
		ip    net.IP
	)
	i, err = net.InterfaceByName(iface)
	if err != nil {
		return
	}
	addrs, err = i.Addrs()
	if err != nil {
		return
	}
	for _, addr := range addrs {
		ip, _, err = net.ParseCIDR(addr.String())
		if err != nil {
			continue
		}
		if ParseIPVersion(ip.String()) == IPv4 {
			ips = append(ips, ip)
		}
	}
	return
}

//convert ip from int64 to string
func InetNtoa(ipnr int64) string {
	var bytes [4]byte
	bytes[0] = byte(ipnr & 0xFF)
	bytes[1] = byte((ipnr >> 8) & 0xFF)
	bytes[2] = byte((ipnr >> 16) & 0xFF)
	bytes[3] = byte((ipnr >> 24) & 0xFF)
	ip := net.IPv4(bytes[3], bytes[2], bytes[1], bytes[0])
	return ip.String()
}

//convert ip from string to int64
func InetAton(ipnr string) int64 {
	bits := strings.Split(ipnr, ".")
	if len(bits) < 4 {
		return 0
	}
	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])
	var sum int64
	sum += int64(b0) << 24
	sum += int64(b1) << 16
	sum += int64(b2) << 8
	sum += int64(b3)
	return sum
}

func GetHostIPs(filter func(ip string) bool) (ips []string) {
	inters, _ := net.Interfaces()
	for _, inter := range inters {
		addrs, _ := inter.Addrs()
		if addrs != nil {
			ipaddr := strings.Split(addrs[0].String(), `/`)
			if filter(ipaddr[0]) {
				ips = append(ips, ipaddr[0])
			}
		}
	}
	return
}
