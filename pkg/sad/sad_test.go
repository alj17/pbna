/* SPDX-License-Identifier: MIT */
package sad

import "testing"

func chkDatabase(db Database) {
}

func TestSAD(t *testing.T) {
	memdb := NewMemDB()
	chkDatabase(memdb)
}
