package builtin

import (
	"fmt"
	"strings"
	"time"

	"github.com/dgurney/mod7/elevendigit"
	"github.com/dgurney/mod7/oem"
	"github.com/dgurney/mod7/tendigit"
	"github.com/dgurney/mod7/validation"
)

/*
   mod7 command of dgtoolbox
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

// generationBenchmark generates 3000000 keys and shows the elapsed time. It's meant to be much more understandable and user-accessible than "make bench"
func generationBenchmark() []string {
	kch := make(chan string)
	keys := make([]string, 0)
	started := time.Now()
	count := 0
	for i := 0; i < 1000000; i++ {
		count++
		go oem.GenerateOEM(kch)
		keys = append(keys, <-kch)
		go tendigit.Generate10digit(kch)
		keys = append(keys, <-kch)
		go elevendigit.Generate11digit(kch)
		keys = append(keys, <-kch)
	}

	fmt.Printf("Took %s to generate %d keys.\n", time.Since(started).Round(time.Millisecond), count*3)
	return keys
}

// validationBenchmark validates 3000000 keys and shows the elapsed time. It's meant to be much more understandable and user-accessible than "make bench"
func validationBenchmark(keys []string) {
	vch := make(chan bool)
	started := time.Now()
	for _, v := range keys {
		go validation.BatchValidate(v, vch)
		<-vch
	}
	fmt.Printf("Took %s to validate %d keys.\n", time.Since(started).Round(time.Millisecond), len(keys))
	return
}

// Mod7 runs the specified part of mod7
func Mod7(part, key string) {
	kch := make(chan string)
	p := strings.ToLower(part)
	switch {
	default:
		fmt.Println("Supported operations: all, bench, cd, oem, 11cd, validate")
		return
	case p == "all":
		go oem.GenerateOEM(kch)
		fmt.Println(<-kch)
		go elevendigit.Generate11digit(kch)
		fmt.Println(<-kch)
		go tendigit.Generate10digit(kch)
		fmt.Println(<-kch)
	case p == "cd":
		go tendigit.Generate10digit(kch)
		fmt.Println(<-kch)
	case p == "oem":
		go oem.GenerateOEM(kch)
		fmt.Println(<-kch)
	case p == "11cd":
		go elevendigit.Generate11digit(kch)
		fmt.Println(<-kch)
	case p == "bench":
		fmt.Println("Running key generation and validation benchmark...")
		validationBenchmark(generationBenchmark())
	case p == "validate":
		validation.ValidateKey(key)
	}

}
