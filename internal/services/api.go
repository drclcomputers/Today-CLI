package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
	"today/internal/logger"
	"today/internal/utils"
)

func getWithRetries(url string, timeout time.Duration, retries int) (*http.Response, error) {
	client := &http.Client{
		Timeout: timeout,
	}

	for i := 0; i < retries; i++ {
		resp, err := client.Get(url)
		if err == nil && resp.StatusCode == http.StatusOK {
			return resp, nil
		}

		if resp != nil {
			resp.Body.Close()
		}

		time.Sleep(500 * time.Millisecond)
	}
	return nil, errors.New("failed to fetch after retries")
}

func fetchAPI(month, day int) utils.ZenquotesTodayAPIResp {
	var resp *http.Response
	var err error
	var url string

	if month == 0 && day == 0 {
		url = utils.TODAY_API_URL
	} else {
		url = fmt.Sprintf(utils.TODAY_API_URL_DATE, month, day)
	}

	resp, err = getWithRetries(url, 5*time.Second, 3)
	if err != nil {
		logger.LogErr(logger.MakeError("API not available. Check your connection or try again later!"))
		return utils.ZenquotesTodayAPIResp{}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	logger.LogErr(err)

	var data utils.ZenquotesTodayAPIResp
	err = json.Unmarshal(body, &data)
	logger.LogErr(err)

	time.Sleep(500 * time.Millisecond)

	return data
}

func GetTodayEventsFromAPI(month, day int) utils.ZenquotesTodayAPIResp {
	return showLoadingSpinner("Fetching data from API...", func() utils.ZenquotesTodayAPIResp {
		return fetchAPI(month, day)
	})
}

func GetEventsType(selection int) []utils.CustomEvent {
	data := GetTodayEventsFromAPI(utils.MONTH, utils.DAY)

	switch selection {
	case 1:
		return data.Data.Events
	case 2:
		return data.Data.Births
	case 3:
		return data.Data.Deaths
	default:
		logger.LogErr(logger.MakeError("Unknown selection! Aborting!"))
	}

	return nil
}
