package req

import (
	"adventofcode/secrets"
	"fmt"
	"io/ioutil"
	"net/http"
)

const aocURL = "https://adventofcode.com/2021"

func MakeRequest(day int) string {
	url := fmt.Sprintf("%s/day/%d/input", aocURL, day)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}

	req.Header.Set("Cookie", secrets.Session)

	client := http.Client{}
	response, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	return string(body)
}
