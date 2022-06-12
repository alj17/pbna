/* SPDX-License-Identifier: MIT */
package sad

import (
	"github.com/alj17/pbna/pkg/inet"
	"github.com/alj17/pbna/pkg/sa"
)

// delete an SA in the SAD
func (db *Database) Delete(inet.IP, sa.SPI, inet.Protocol, inet.Family, sa.Mark) error {
	return nil
}
