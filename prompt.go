package main

/*
   The prompt-related parts of dgtoolbox
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
	"os"
)

func ps1() string {
	dir, err := os.Getwd()
	if err != nil {
		dir = "broken"
	}
	symbol := " $"

	return dir + symbol
}
