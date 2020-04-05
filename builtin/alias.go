package builtin

import (
	"fmt"
	"strings"
)

/*
   alias functionality of dgtoolbox
   Copyright (C) 2019 Daniel Gurney
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

var aliases map[string]string

// Alias allows the user to define an alias.
func Alias(a string, c []string) {
	if aliases == nil {
		aliases = make(map[string]string)
	}
	aliases[a] = strings.Join(c, " ")
}

// Aliases prints all currently defined aliases.
func Aliases() {
	for alias, command := range aliases {
		fmt.Println(alias, "=", command)
	}
	if len(aliases) == 0 {
		fmt.Println("No defined aliases.")
	}
}

// DeleteAlias deletes the specified alias.
func DeleteAlias(a string) {
	delete(aliases, a)
}

// GetAliasCommand returns the provided alias command if it exists. If not, the wannabe alias is rudely ignored, and remains unloved.
func GetAliasCommand(a string) string {
	if _, exists := aliases[a]; !exists {
		return ""
	}
	return aliases[a]
}
