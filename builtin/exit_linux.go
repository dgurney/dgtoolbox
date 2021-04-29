package builtin

/*
   Exit command of dgtoolbox. For other commands, see the builtin/ directory.
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
	"syscall"
)

// Exit exits the program.
func Exit() {
	switch {
	default:
		os.Exit(0)
	case os.Getpid() == 1:
		syscall.Reboot(syscall.LINUX_REBOOT_CMD_POWER_OFF)
	}
}
