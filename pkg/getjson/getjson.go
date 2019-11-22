package getjson

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type RetrieveJson func(fullUrl string, client *http.Client) map[string]interface{}

func Map(url string, client *http.Client) map[string]interface{} {
	var result map[string]interface{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	jsonErr := json.Unmarshal(body, &result)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return result
}