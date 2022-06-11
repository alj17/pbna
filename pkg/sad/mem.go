/* SPDX-License-Identifier: MIT */
package sad

import (
	"github.com/alj17/pbna/pkg/inet"
	"github.com/alj17/pbna/pkg/sa"
)

type MemDB struct {
}

func NewMemDB() *MemDB {
	return nil
}

func (mem *MemDB) New(sa *sa.SecurityAssociation) error {
	return nil
}

func (mem *MemDB) Delete(inet.IP, sa.SPI, inet.Protocol, inet.Family, sa.Mark) error {
	return nil
}

func (mem *MemDB) Get(inet.IP, sa.SPI, inet.Protocol, inet.Family, sa.Mark) error {
	return nil
}
