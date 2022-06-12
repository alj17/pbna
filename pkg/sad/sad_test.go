/* SPDX-License-Identifier: MIT */
package sad

import "testing"

type notifier struct {
	serial uint64
	t *testing.T
}

func (n notifier) Notify(db *Database) {
	chgs, err := db.Changes(n.serial)

	if err != nil {
		n.t.Fatalf("unable to fetch changes\n")
	}

	for _, chg := range chgs {
		// sync change to backend
		n.serial = chg.Serial
	}

	db.AckChanges(n.serial)
}

func TestSAD(t *testing.T) {
	dbn := notifier{t: t}
	db := NewDatabaseNotify(dbn)
	db.Sync()

	// db.Add()
	// db.Get()
	// db.Update()
	// db.Sync()
	// db.Delete()
	// db.Sync()
	// db.Clear()
	// db.Sync()
}
