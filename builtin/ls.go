package builtin

import (
	"fmt"
	"path/filepath"
)

/*
   ls command of dgtoolbox
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

// Ls does a very rudimentary directory listing
func Ls() {
	listing, err := filepath.Glob("*")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, l := range listing {
		fmt.Println(l)
	}
}
