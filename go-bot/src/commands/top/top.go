package top

import (
	"bot"
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"strconv"
)

type esResult struct {
	Success bool `json:success`
	Reason string `json:reason`
	Result []struct {
		Nick string `json:nick`
		Delay float64 `json:delay`
	} `json:result`
}

func top(command *bot.Cmd) (msg string, err error) {
	var apiUrl = command.APIEndpoint +"/top"

	client := &http.Client{}
	req, _ := http.NewRequest("GET", apiUrl, nil)

	q := req.URL.Query()

	if len(command.Args) >= 1 {
		q.Add("date", command.Args[0])
	}
	q.Add("channel", command.Channel)
	q.Add("EndpointKey", command.APIEndpointKey)

	req.URL.RawQuery = q.Encode()
	res, err := client.Do(req)

	body, err := ioutil.ReadAll(res.Body)

	var leets esResult
	esErr := json.Unmarshal(body, &leets)
	if esErr != nil {
		fmt.Println(esErr)
	}

	if leets.Success == false {
		return command.User.Nick +": "+ leets.Reason, nil
	}

	var str = command.User.Nick +": "

	for i, a := range leets.Result {
		var separator = " - "

		if i+1 == len(leets.Result) {
			separator = ""
		}

		str += a.Nick +" "+ strconv.FormatFloat(a.Delay, 'f', -1, 64) + separator
	}


	return str, nil
}

func init() {
	bot.RegisterCommand(
		"top",
		"Returns top 3 leets for a channel on any given date. Defaults to today",
		"",
		top,
		false,
		false)
}
