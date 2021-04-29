package builtin

import (
	"fmt"
	"strings"

	"github.com/dgurney/unikey/generator"
)

/*
   chicagokey command of dgtoolbox
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

// Chicagokey generates a site id and password for the specified build
func Chicagokey(build string) {
	b := strings.ToLower(build)
	switch {
	default:
		fmt.Println("Supported builds: 73f/73g/81 (up to 90c)/99 (up to 116)/122 (up to 189)/216 (up to 302)")
		return
	case b == "73f" || b == "73g" || b == "81" || b == "99" || b == "122" || b == "216":
		k, _ := generator.Generate(generator.ChicagoCredentials{Build: b})
		fmt.Printf("%s: %s\n", b, k.String())
	}
}
