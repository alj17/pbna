/* SPDX-License-Identifier: MIT */
package main

import "fmt"

func xfrm(args []string) {
	cmd := args[0]
	if cmd == "state" {
		xfrm_state(args[1:])
	} else if cmd == "policy" {
		xfrm_policy(args[1:])
	} else if cmd == "monitor" {
		xfrm_monitor(args[1:])
	} else {
		xfrm_usage()
	}
}

func xfrm_usage() {
	fmt.Println("Usage: pbna xfrm state|policy|monitor ...")
}
