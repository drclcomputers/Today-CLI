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
