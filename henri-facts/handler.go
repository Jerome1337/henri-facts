package function

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/openfaas-incubator/go-function-sdk"
)

type facts struct {
	Facts []string `json:"facts"`
}

type slackResponse struct {
	ResponseType string `json:"response_type"`
	Text         string `json:"text"`
}

func Handle(req handler.Request) (handler.Response, error) {
	var err error

	rand.Seed(time.Now().UTC().UnixNano())

	jsonFile, jsonErr := os.Open("facts.json")

	if jsonErr != nil {
		return handler.Response{
			Body:       []byte(fmt.Sprintf("Failed to open JSON file: %s", jsonErr)),
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	defer jsonFile.Close()

	byteArray, _ := ioutil.ReadAll(jsonFile)

	var facts facts

	if err := json.Unmarshal(byteArray, &facts); err != nil {
		return handler.Response{
			Body:       []byte(fmt.Sprintf("Failed to unmarshal JSON: %s", err)),
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	marshalJSON, _ := json.Marshal(
		slackResponse{
			"in_channel",
			facts.Facts[randomInt(0, len(facts.Facts))],
		},
	)

	return handler.Response{
		Body: marshalJSON,
		Header: map[string][]string{
			"Content-Type": []string{"application/json"},
		},
		StatusCode: http.StatusOK,
	}, err
}

func randomInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
