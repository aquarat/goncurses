// Copyright 2011 Rob Thornton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

import gc "code.google.com/p/goncurses"

func main() {
	stdscr, err := gc.Init()
	if err != nil {
		return
	}
	defer gc.End()

	stdscr.Border(gc.VLINE, gc.ACS_ACS_VLINE, gc.ACS_HLINE, gc.ACS_HLINE,
		gc.ACS_ULCORNER, gc.ACS_LRCORNER, gc.ACS_ULCORNER, gc.ACS_LRCORNER)
	stdscr.GetChar()
}
