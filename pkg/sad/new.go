/* SPDX-License-Identifier: MIT */
package sad

import (
// "github.com/alj17/pbna/pkg/inet"
// "github.com/alj17/pbna/pkg/sa"
)

func NewDatabaseNotify(dbn Notify) *Database {
	return nil
}

func NewDatabase() *Database {
	nn := notify{}
	return NewDatabaseNotify(nn)
}

// ************* Implementation (private) *******************

// minimal notifier
type notify struct {
	serial uint64
}

func (n notify) Notify(db *Database) {
	n.serial = db.Serial()
	db.AckChanges(n.serial)
}
