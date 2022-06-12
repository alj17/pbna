/* SPDX-License-Identifier: MIT */
package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		usage()
	}

	cmd := args[1]
	if cmd == "state" {
		state(args[2:])
	} else if cmd == "policy" {
		policy(args[2:])
	} else if cmd == "monitor" {
		monitor(args[2:])
	} else {
		usage()
	}
}

func usage() {
	fmt.Println("Usage: xfrm state|policy|monitor ...")
	os.Exit(1)
}
