package main

/*
   dgtoolbox - a lackluster busybox-like program
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
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/dgurney/dgtoolbox/builtin"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	ver()
	fmt.Println(`Run "help" to see the list of built-in commands and shortcuts.`)
	fmt.Println(runtime.NumCPU(), "CPUs available.")
	for {
	Beginning:
		fmt.Print(ps1() + " ")
		i, _ := reader.ReadString('\n')
		// Ctrl+d
		if len(i) == 0 {
			i = "exit"
			fmt.Println(i)
		}
		if i == "\n" || i == "\r\n" {
			goto Beginning
		}
		input := strings.Fields(strings.TrimSuffix(i, "\n"))

		// Handle aliases first, otherwise we could not alias built-in commands without convoluted if conditions
		if len(builtin.GetAliasCommand(input[0])) != 0 {
			// since our shell is naive, we must split the returned command before running it
			ac := strings.Split(builtin.GetAliasCommand(input[0]), " ")
			input = ac
		}

		switch input[0] {
		case "~":
			builtin.Shortcut(input[0])
		case "..":
			builtin.Shortcut(input[0])
		case "exit":
			builtin.Exit()
		case "cd":
			if len(input) == 1 {
				fmt.Println("Specify a directory")
				goto Beginning
			}
			builtin.Cd(input[1])
		case "ver":
			ver()
		case "help":
			builtin.Help()
		case "echo":
			if len(input) == 1 {
				fmt.Println("")
				goto Beginning
			}
			builtin.Echo(input[1:])
		case "alias":
			switch {
			default:
				builtin.Alias(input[1], input[2:])
			case len(input) == 1:
				builtin.Aliases()
			case len(input) == 2:
				builtin.DeleteAlias(input[1])
			}
		case "mkdir":
			builtin.Mkdir(input[1])
		case "ls":
			builtin.Ls()
		case "mod7":
			switch {
			default:
				builtin.Mod7("", "")
			case len(input) == 2:
				builtin.Mod7(input[1], "")
			case len(input) == 3:
				builtin.Mod7(input[1], input[2])
			}
		case "chicagokey":
			switch {
			default:
				builtin.Chicagokey("")
			case len(input) == 2:
				builtin.Chicagokey(input[1])
			}
		default:
			builtin.Help()
		}
	}
}
