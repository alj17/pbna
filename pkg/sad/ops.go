/* SPDX-License-Identifier: MIT */
package sad

import (
	"github.com/alj17/pbna/pkg/inet"
	"github.com/alj17/pbna/pkg/sa"
)

type Database interface {
	New(*sa.SecurityAssociation) error
	Delete(inet.IP, sa.SPI, inet.Protocol, inet.Family, sa.Mark) error
	Get(inet.IP, sa.SPI, inet.Protocol, inet.Family, sa.Mark) error
}
