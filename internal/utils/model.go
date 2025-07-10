// Copyright (c) 2025 @drclcomputers. All rights reserved.
//
// This work is licensed under the terms of the MIT license.
// For a copy, see <https://opensource.org/licenses/MIT>.

package utils

type ZenquotesTodayAPIResp struct {
	Date string `json:"date"`
	Data Data   `json:"data"`
}

type CustomEvent struct {
	Text  string `json:"text"`
	Links struct {
		Info struct {
			Year string `json:"2"`
		} `json:"0"`
	} `json:"links"`
}

type Data struct {
	Events []CustomEvent `json:"Events"`
	Births []CustomEvent `json:"Births"`
	Deaths []CustomEvent `json:"Deaths"`
}
