package builtin

import (
	"fmt"
	"strings"
	"time"

	"github.com/dgurney/unikey/generator"
	"github.com/dgurney/unikey/validator"
)

/*
   mod7 command of dgtoolbox
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

func generationBenchmark(amount int) []string {
	oem := generator.Mod7OEM{}
	cd := generator.Mod7CD{}
	ecd := generator.Mod7ElevenCD{}
	keys := make([]string, 0)
	started := time.Now()
	for i := 0; i < amount; i++ {
		k, _ := generator.Generate(oem)
		keys = append(keys, k.String())
	}
	for i := 0; i < amount; i++ {
		k, _ := generator.Generate(cd)
		keys = append(keys, k.String())
	}
	for i := 0; i < amount; i++ {
		k, _ := generator.Generate(ecd)
		keys = append(keys, k.String())
	}

	var ended time.Duration
	switch {
	case time.Since(started).Round(time.Second) > 1:
		ended = time.Since(started).Round(time.Millisecond)
	default:
		ended = time.Since(started).Round(time.Microsecond)
	}
	fmt.Printf("Took %s to generate %d keys.\n", ended, len(keys))

	return keys
}

func validationBenchmark(keys []string) {
	var ki validator.KeyValidator
	started := time.Now()
	for _, k := range keys {
		switch {
		case len(k) == 12 && k[4:5] == "-":
			ki = validator.Mod7ElevenCD{
				First:  k[0:4],
				Second: k[5:12],
			}
		case len(k) == 11 && k[3:4] == "-":
			ki = validator.Mod7CD{
				First:  k[0:3],
				Second: k[4:11],
			}
		case len(k) == 23 && k[5:6] == "-" && k[9:10] == "-" && k[17:18] == "-" && len(k[18:]) == 5:
			ki = validator.Mod7OEM{
				First: k[0:5],
				// nice
				Second: k[6:9],
				Third:  k[10:17],
				Fourth: k[18:],
			}
		}
		validator.Validate(ki)
	}

	var ended time.Duration
	switch {
	case time.Since(started).Round(time.Second) > 1:
		ended = time.Since(started).Round(time.Millisecond)
	default:
		ended = time.Since(started).Round(time.Microsecond)
	}

	fmt.Printf("Took %s to validate %d keys.\n", ended, len(keys))
}

// Mod7 runs the specified part of mod7
func Mod7(part, key string, is95 bool) {
	p := strings.ToLower(part)
	switch {
	default:
		fmt.Println("Supported operations: all, bench, cd, oem, 11cd, validate")
		return
	case p == "all":
		oemkey, _ := generator.Generate(generator.Mod7OEM{})
		ecdkey, _ := generator.Generate(generator.Mod7ElevenCD{})
		cdkey, _ := generator.Generate(generator.Mod7CD{})
		fmt.Println(oemkey.String())
		fmt.Println(ecdkey.String())
		fmt.Println(cdkey.String())
	case p == "cd":
		cdkey, _ := generator.Generate(generator.Mod7CD{})
		fmt.Println(cdkey.String())
	case p == "oem":
		oemkey, _ := generator.Generate(generator.Mod7OEM{})
		fmt.Println(oemkey.String())
	case p == "11cd":
		ecdkey, _ := generator.Generate(generator.Mod7ElevenCD{})
		fmt.Println(ecdkey.String())
	case p == "bench":
		keys := generationBenchmark(100000)
		validationBenchmark(keys)
	case p == "validate":
		var ki validator.KeyValidator

		switch {
		case len(key) == 12 && key[4:5] == "-":
			ki = validator.Mod7ElevenCD{
				First:  key[0:4],
				Second: key[5:12],
			}
		case len(key) == 11 && key[3:4] == "-":
			ki = validator.Mod7CD{
				First:  key[0:3],
				Second: key[4:11],
				Is95:   is95,
			}
		case len(key) == 23 && key[5:6] == "-" && key[9:10] == "-" && key[17:18] == "-" && len(key[18:]) == 5:
			ki = validator.Mod7OEM{
				First: key[0:5],
				// nice
				Second: key[6:9],
				Third:  key[10:17],
				Fourth: key[18:],
				Is95:   is95,
			}
		default:
			fmt.Println("Could not recognize key type")
			return
		}
		err := validator.Validate(ki)
		switch {
		default:
			fmt.Printf("%s is valid\n", key)
		case err != nil:
			fmt.Printf("%s is invalid: %s\n", key, err)
		}
	}

}
