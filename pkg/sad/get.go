/* SPDX-License-Identifier: MIT */
package sad

import (
	"github.com/alj17/pbna/pkg/inet"
	"github.com/alj17/pbna/pkg/sa"
)

// lookup an SA in the SAD
func (db *Database) Get(inet.IP, sa.SPI, inet.Protocol, inet.Family, sa.Mark) error {
	return nil
}

