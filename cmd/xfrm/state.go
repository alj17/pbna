/* SPDX-License-Identifier: MIT */
package main

import (
	"fmt"
	"github.com/alj17/pbna/pkg/sad"
)

func state(args []string) {
	var cmd string
	//var dbn sad.Notify
	//var notify bool

	// pull off options first
	for _, arg := range args {
		if arg[0] != '-' {
			break
		}

		if len(arg) > 0 && arg[1] == 'n' {
			if args[0] == "xfrm" { // create an xfrm
				// dbn = xfrm.StateNotify()
				// notify = true
			} else {
				state_usage()
				return
			}
		} else {
			state_usage()
			return
		}
	}

	db := sad.NewDatabase()

	if args == nil || len(args) == 0 {
		cmd = "list"
	} else {
		cmd = args[0]
		args = args[1:]
	}

	if cmd == "add" || cmd == "update" {
		state_add_update(db, cmd, args)
	} else if cmd == "allocspi" {
		state_allocspi(db, args)
	} else if cmd == "delete" || cmd == "get" {
		state_delete_get(db, cmd, args)
	} else if cmd == "deleteall" || cmd == "list" || cmd == "ls" {
		state_deleteall_list(db, cmd, args)
	} else if cmd == "flush" {
		state_flush(db, args)
	} else if cmd == "count" && len(args) == 0 {
		state_count(db)
	} else {
		state_usage()
	}
}

func state_usage() {
	fmt.Println("Usage: xfrm state add|update|allocspi|delete|get|list|flush|count")
}

func state_add_update(db *sad.Database, cmd string, args []string) {
}

func state_allocspi(db *sad.Database, args []string) {
}

func state_delete_get(db *sad.Database, cmd string, args []string) {
}

func state_deleteall_list(db *sad.Database, cmd string, args []string) {
}

func state_flush(db *sad.Database, args []string) {
}

func state_count(db *sad.Database) {
}
