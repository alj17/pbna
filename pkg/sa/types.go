/* SPDX-License-Identifier: MIT */
package sa

import (
	"github.com/alj17/pbna/pkg/inet"
)

type SPI uint32
type Mode int

const (
	TRANSPORT Mode = iota
	TUNNEL
	ROUTEOPTIMIZATION
	IN_TRIGGER
	BEET
)

type Mark struct {
	value uint32
	mask  uint32
}

type RequestID uint32

type SecurityAssociation struct {
	Source      inet.IP
	Destination inet.IP
	SPI         SPI
	Mode        Mode
	Protocol    inet.Protocol

	Mark      Mark
	RequestID RequestID
}
