/* SPDX-License-Identifier: MIT */
package inet

import "net"

type IP struct {
	addr uint32
}

func IPv4(a, b, c, d uint8) IP {
	var ip uint32

	ip = (uint32(a) << 24) |
		(uint32(b) << 16) |
		(uint32(c) << 8) |
		(uint32(d))

	return IP{addr: ip}
}

func (ip IP) NetIP() net.IP {
	var a, b, c, d uint8

	a = uint8(ip.addr >> 24)
	b = uint8(ip.addr >> 16)
	c = uint8(ip.addr >> 8)
	d = uint8(ip.addr)

	return net.IPv4(a, b, c, d)
}

func (ip IP) String() string {
	return ip.NetIP().String()
}

func NetIP(ip net.IP) IP {
	ip = ip.To4()
	// FIXME: ip == nil?
	return IPv4(ip[0], ip[1], ip[2], ip[3])
}

func ParseIP(s string) IP {
	return NetIP(net.ParseIP(s))
}
