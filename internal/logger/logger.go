// Copyright (c) 2025 @drclcomputers. All rights reserved.
//
// This work is licensed under the terms of the MIT license.
// For a copy, see <https://opensource.org/licenses/MIT>.

package logger

import (
	"fmt"
	"log"

	"github.com/fatih/color"
)

func LogErr(err error) {
	if err != nil {
		log.Fatal(color.New(color.FgWhite, color.BgRed).Sprint(err))
	}
}

func MakeError(text string) error {
	return fmt.Errorf("%s", text)
}
