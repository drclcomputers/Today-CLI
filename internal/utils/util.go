// Copyright (c) 2025 @drclcomputers. All rights reserved.
//
// This work is licensed under the terms of the MIT license.
// For a copy, see <https://opensource.org/licenses/MIT>.

package utils

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"

	"golang.design/x/clipboard"
	"golang.org/x/term"
)

const TODAY_API_URL = "https://today.zenquotes.io/api"

const TODAY_API_URL_DATE = TODAY_API_URL + "/%d/%d" // month, day

var _, month, day = time.Now().Date()
var MONTH int = int(month)
var DAY int = day

func GetTerminalSize() (int, int) {
	width, height, err := term.GetSize(0)
	if err != nil {
		return 80, 20
	}
	return width, height
}

func IntToMonthName(month int) string {
	if month < 1 || month > 12 {
		return "Invalid"
	}
	return time.Month(month).String()
}

func ClearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		fmt.Printf("Warning: Failed to clear screen: %v\n", err)
	}
}

func CopyToClipboard(text string) {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}
	clipboard.Write(clipboard.FmtText, []byte(text))
}
