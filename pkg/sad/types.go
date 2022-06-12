/* SPDX-License-Identifier: MIT */
package sad

import (
	"github.com/alj17/pbna/pkg/sa"
)

// Security Association Database
// see:
//    new.go to create a Database
//    add.go to add an SA to the Database
//    update.go to update an entry in the Database
//    get.go to retrieve an entry from the database
//    delete.go to delete an entry from the database
//    changes.go to retrieve changes since serial in the database
type Database struct {
	serial	uint64		// serial number changes on every update
}

// Database Notifications -- will be called with serial == 0 on startup and
// at "appropriate" times throughout the DB lifecycle.  Call db.Sync() to
// sync database, see changes.go for details
type Notify interface {
	Notify(*Database)
}

// Filter for db.Walk() -- see walk.go
type Filter struct {
	spi	sa.SPI
}

// Callback for db.Walk() -- see walk.go
type Callback interface {
	Entry(*Database, *sa.SecurityAssociation)
}

// change (operation) to DB -- see changes.go
type Change struct {
	Serial	uint64	// serial number of this change
}
