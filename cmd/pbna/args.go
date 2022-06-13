/* SPDX-License-Identifier: MIT */
package main

import (
	"github.com/alj17/pbna/pkg/inet"
	"github.com/alj17/pbna/pkg/sa"
	"strconv"
)

type identifier struct {
	src   inet.IP
	dst   inet.IP
	proto inet.Protocol
	spi   sa.SPI
}

func parseID(args []string) ([]string, *identifier, error) {
	id := identifier{}
	for len(args) > 0 {
		if args[0] == "src" && len(args) > 1 {
			id.src = inet.ParseIP(args[1])
			args = args[1:]
		} else if args[0] == "dst" && len(args) > 1 {
			id.dst = inet.ParseIP(args[1])
			args = args[1:]
		} else if args[0] == "proto" && len(args) > 1 {
			if args[1] == "esp" {
				id.proto = inet.ESP
			} else if args[1] == "ah" {
				id.proto = inet.AH
			} else if args[1] == "comp" {
				id.proto = inet.COMP
			} else {
				args = args[1:]
				// FIXME: error!
				return args, &id, nil
			}

			args = args[1:]
		} else if args[0] == "spi" && len(args) > 1 {
			i, err := strconv.ParseInt(args[1], 0, 32)
			if err != nil {
				return args[2:], &id, err
			}
			id.spi = sa.SPI(i)
			args = args[1:]
		}

		if len(args) > 0 {
			args = args[1:]
		}
	}
	return args, &id, nil
}
