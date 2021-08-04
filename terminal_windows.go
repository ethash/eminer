// +build windows

package main

import (
	"os"

	"golang.org/x/sys/windows"
)

func init() {
	var outMode uint32
	out := windows.Handle(os.Stdout.Fd())

	if err := windows.GetConsoleMode(out, &outMode); err == nil {
		if err := windows.SetConsoleMode(out, outMode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING); err != nil {
			windows.SetConsoleMode(out, outMode)
			return
		}
	} else {
		return
	}

	virtualTerm = true
}
