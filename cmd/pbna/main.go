/* SPDX-License-Identifier: MIT */
package main

import (
	"fmt"
	"os"
)

func main() {
	var batch string

	args := os.Args[1:]
	for len(args) > 1 {
		if args[0] == "-b" {
			batch = args[1]
			args = args[1:]
		} else {
			break
		}
		args = args[1:]
	}

	if batch != "" {
		return
	}

	if len(args) < 2 {
		usage()
	}

	if args[0] == "xfrm" {
		xfrm(args[1:])
	}
}

func usage() {
	fmt.Println("Usage: pbna [options] xfrm|link|addr|trace ...")
	fmt.Println("  options:")
	fmt.Println("     -b file    # take commands from file")
	os.Exit(1)
}
