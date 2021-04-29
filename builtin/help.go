package builtin

/*
   help command of dgtoolbox
   Copyright (C) 2021 Daniel Gurney
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
)

// Help prints all currently supported commands and shortcuts.
func Help() {
	fmt.Println("The following built-in commands are available:")
	commands := []string{"alias", "cd", "chicagokey", "echo", "exit", "help", "ls", "mkdir", "mod7", "ver"}
	for _, c := range commands {
		fmt.Println(c)
	}
	fmt.Println("The following built-in shortcuts are available:")
	shortcuts := []string{".."}
	for _, s := range shortcuts {
		fmt.Println(s)
	}
}
