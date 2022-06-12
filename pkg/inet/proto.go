/* SPDX-License-Identifier: MIT */
package inet

type Protocol uint8

const (
	ICMP Protocol = 1
	IPIP          = 4
	TCP           = 6
	UDP           = 17
	ESP           = 50
	AH            = 51
	COMP          = 108
)
