/* SPDX-License-Identifier: MIT */
package inet

import "testing"

func TestIP(t *testing.T) {
	ip := IPv4(1, 2, 3, 4)

	if ip.String() != "1.2.3.4" {
		t.Fatalf("%s != %s", ip, "1.2.3.4")
	}
}
