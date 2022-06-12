/* SPDX-License-Identifier: MIT */
package sad

import (
	//"github.com/alj17/pbna/pkg/inet"
	//"github.com/alj17/pbna/pkg/sa"
)

// return a list of changes to database since serial number "from"
//
// len(changes) == 0 if from == 0
//
func (db *Database) Changes(from uint64) ([] Change, error) {
	return nil, nil
}

func (db *Database) Serial() uint64 {
	return db.serial
}

func (db *Database) AckChanges(serial uint64) error {
	return nil
}

// ensure all changes are sent to notifier
func (db *Database) Sync() error {
	return nil
}
