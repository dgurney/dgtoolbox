package main

/*
   The most fundamental commands of dgtoolbox. For other commands, see the builtin/ directory.
   Copyright (C) 2020 Daniel Gurney
   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.
   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.
   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

import (
	"fmt"
	"os"
)

func exit() {
	os.Exit(0)
}

func ver() {
	fmt.Printf("dgtoolbox v%s by Daniel Gurney. Licensed under GPLv3.\n", getVersion())
}
